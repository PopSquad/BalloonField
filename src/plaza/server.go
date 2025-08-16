package plaza

import (
	"net"
	"time"

	"github.com/PopSquad/BalloonField/src/network"
	"github.com/PopSquad/BalloonField/src/util"
)

type Server struct {
	Address string
	Logger  util.Logger
}

func NewPlazaServer(address string) *Server {
	return &Server{
		Address: address,
		Logger: util.Logger{
			Tag: "PlazaServer",
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
		client := NewPlazaClient(conn)
		go client.Loop()
	}
}

type Client struct {
	*network.PPConn
	Logger util.Logger
}

func NewPlazaClient(conn net.Conn) *Client {
	return &Client{
		PPConn: network.NewPPConn(conn),
		Logger: util.Logger{
			Tag: "PlazaClient",
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
		if payload[0] == CMD96 && payload[1] == PRM96_01 {
			err = t.Write([]byte{CMD96, PRM96_64})
			if err != nil {
				break
			}
			err = t.Write([]byte{CMD97, PRM97_64}, BuildParam97_64Data(0x00000001, []SvrNodeInfo{
				MockRecord(),
			}))
			if err != nil {
				break
			}
			err = t.Write([]byte{CMD97, PRM97_67, 0xAB, 0xAB})
			if err != nil {
				break
			}
		}
		if payload[0] == CMD97 && payload[1] == PRM97_00 {
			break
		}
	}

	t.Logger.Log("%s leave", t.ConnAddr)
}
