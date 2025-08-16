package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/PopSquad/BalloonField/src/cipher"
	"github.com/PopSquad/BalloonField/src/util"
)

func NewPPConn(conn net.Conn) *PPConn {
	return &PPConn{
		ConnID:   util.GenID(),
		ConnAddr: conn.RemoteAddr().String(),
		Conn:     conn,
		Decipher: cipher.Decipher{},
		Cipher:   cipher.Cipher{},
		Logger: util.Logger{
			Tag: "PPConn",
		},
	}
}

type PPConn struct {
	ConnID   int64
	ConnAddr string
	Conn     net.Conn
	Cipher   cipher.Cipher
	Decipher cipher.Decipher
	Logger   util.Logger
}

func (t *PPConn) writePkt(data []byte) (err error) {
	sizeBuf := make([]byte, 2)
	binary.LittleEndian.PutUint16(sizeBuf, uint16(len(data)))
	_, err = t.Conn.Write(sizeBuf)
	if err != nil {
		return
	}
	_, err = t.Conn.Write(data)
	return
}

func (t *PPConn) readPkt() (pktData []byte, err error) {
	sizeBuf := make([]byte, 2)
	_, err = io.ReadFull(t.Conn, sizeBuf)
	if err != nil {
		return
	}
	pktSize := binary.LittleEndian.Uint16(sizeBuf)
	if pktSize <= 0 || pktSize > 16640 {
		err = fmt.Errorf("invalid pktSize %d", pktSize)
	}
	pktData = make([]byte, pktSize)
	_, err = io.ReadFull(t.Conn, pktData)
	return
}

func (t *PPConn) Write(payload ...[]byte) (err error) {
	d, err := t.Cipher.Encrypt(bytes.Join(payload, nil))
	if err != nil {
		t.Logger.Log("failed to encrypt: %v", err)
		return
	}
	err = t.writePkt(d)
	if err != nil {
		t.Logger.Log("failed to write: %v", err)
		return
	}
	t.Logger.Log("resp: %x", bytes.Join(payload, nil))
	return
}

func (t *PPConn) Read() (payload []byte, err error) {
	pkt, err := t.readPkt()
	if err != nil {
		t.Logger.Log("failed to read: %v", err)
		return
	}
	payload, err = t.Decipher.Decrypt(pkt)
	if err != nil {
		t.Logger.Log("failed to decrypt: %v", err)
		return
	}
	t.Logger.Log("payload(%dB): %x", len(payload), payload)
	return
}
