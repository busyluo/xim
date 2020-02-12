package codec

import (
	"encoding/binary"
	"errors"
	"net"
)

// LV结构中, L所占的长度
var LenFieldLen = 3

type LvCodec struct {
	net.Conn
}

func NewLvCodec(conn net.Conn) *LvCodec {
	return &LvCodec{conn}
}

func (c *LvCodec) Decode() ([]byte, error) {
	var header = make([]byte, 4)
	c.Conn.Read(header)
	if header[0] != 0xAA {
		return nil, errors.New("illegal frame")
	}
	length := int(binary.BigEndian.Uint32(header[1:]))

	var data = make([]byte, length)
	c.Conn.Read(data)
	return data, nil
}

func (c *LvCodec) Encode(data []byte) {
	var l = len(data)
	var buffer = make([]byte, LenFieldLen+len(data))
	binary.BigEndian.PutUint32(buffer[0:4], uint32(l))
	buffer[0] = 0xAA

	copy(buffer[LenFieldLen:], data)

	c.Conn.Write(buffer)
}
