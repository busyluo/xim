package framedserver

import (
	"encoding/binary"
	"errors"
	"net"
)

type FramedReader interface {
	ReadFrame()
}

type FramedConn struct {
	*net.TCPConn
}

func (conn *FramedConn) ReadFrame() ([]byte, error) {
	var header = make([]byte, 4)
	conn.Read(header)
	if header[0] != 0xAA {
		return nil, errors.New("illegal frame")
	}
	length := int(binary.BigEndian.Uint32(header[1:]))

	var data = make([]byte, length)
	conn.Read(data)
	return data, nil
}
