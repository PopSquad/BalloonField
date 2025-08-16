package room

import (
	"bytes"
	"net"
	"time"

	"github.com/PopSquad/BalloonField/src/network"
	"github.com/PopSquad/BalloonField/src/util"
)

type Server struct {
	Address string
	Logger  util.Logger
}

func NewRoomServer(address string) *Server {
	return &Server{
		Address: address,
		Logger: util.Logger{
			Tag: "RoomServer",
		},
	}
}

func (t *Server) Start() {
	listener, err := net.Listen("tcp", t.Address)
	if err != nil {
		t.Logger.Log("failed to listen: %v", err)
		panic("oops")
	}
	defer listener.Close()

	t.Logger.Log("running at %s", t.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Logger.Log("failed to accept client: %v", err)
			continue
		}
		client := NewRoomClient(conn)
		go client.Loop()
	}
}

type Client struct {
	*network.PPConn
	Logger util.Logger
}

func NewRoomClient(conn net.Conn) *Client {
	return &Client{
		PPConn: network.NewPPConn(conn),
		Logger: util.Logger{
			Tag: "RoomClient",
		},
	}
}

func (t *Client) Loop() {
	t.Conn.SetDeadline(time.Now().Add(1000 * time.Second))
	defer t.Conn.Close()

	t.Logger.Log("%s incoming", t.ConnAddr)

	for {
		payload, err := t.Read()
		if err != nil {
			break
		}
		if bytes.HasPrefix(payload, []byte{CMD00, PRM00_02}) {
			err = t.Write([]byte{CMD00, PRM00_64}, BuildLoginRespPacket(LoginRespPacket{
				ResultCode: DAT00_64_00,
				Unk1:       0,
				Unk2:       0,
				Unk3:       [4]byte{},
				Unk4:       [17]byte{},
				MsgLen:     0,
				Msg:        "",
			}))
			if err != nil {
				break
			}
			err = t.Write([]byte{CMD01, PRM01_6A, 0x01, 0x01}) // TODO: not working yet
			if err != nil {
				break
			}
		}
	}

	t.Logger.Log("%s leave", t.ConnAddr)
}
