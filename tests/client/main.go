package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"time"
	"xim/internal/tcpserver/codec"
	"xim/internal/tcpserver/pb"
)

type TcpClient struct {
	pb.App
	DeviceId int64
	UserId   int64
	Token    string

	codec *codec.LvCodec
}

func (c *TcpClient) Login() {
	req := &pb.LoginReq{
		App:      c.App,
		DeviceId: c.DeviceId,
		UserId:   c.UserId,
		Token:    c.Token,
	}
	id := pb.RequestId(0)
	buf := pb.MarshalRequest(id, req)

	c.codec.Encode(buf)
}

func (c *TcpClient) Start() error {
	conn, err := net.Dial("tcp", "127.0.0.1:6060")
	if err != nil {
		return err
	}
	c.codec = codec.NewLvCodec(conn)
	return nil
}

func (c *TcpClient) Heartbeat() {
	ticker := time.NewTicker(time.Minute * 5)
	for range ticker.C {
		//
	}
}

func (c *TcpClient) Receive() {
	for {
		bytes, err := c.codec.Decode()
		if err != nil {
			fmt.Println(err)
			return
		}
		c.HandlePacket(bytes)
	}
}

func (c *TcpClient) HandlePacket(data []byte) {
	var packet pb.ResponsePacket
	err := proto.Unmarshal(data, &packet)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch packet.Type {
	case pb.PacketType_LOGIN:
		fmt.Println("Login ok.")
	case pb.PacketType_HEARTBEAT:
		fmt.Println("Heart Beat Ack.")
	}
}

func main() {
	client := TcpClient{
		App:      pb.App_LINUX,
		DeviceId: 0,
		UserId:   0,
		Token:    "",
	}
	err := client.Start()
	if err != nil {
		fmt.Println("Start failed.")
		return
	}
	go client.Receive()

	client.Login()
}
