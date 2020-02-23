package logic

import (
	"github.com/golang/protobuf/proto"
	tcppb "xim/internal/app/tcpserver/pb"
	"xim/internal/pkg/pb"
)

func DispatchMessage(devices []int64, m *pb.DownMessage) error {
	for _, d := range devices {
		conn := manager.Load(d)

		data, _ := proto.Marshal(m)
		dispatchData(conn, tcppb.PacketType_PACKET_MESSAGE, data)
	}
	return nil
}

func dispatchData(conn *ConnCtx, t tcppb.PacketType, data []byte) {
	res := tcppb.RequestPacket{
		Type: t,
		Data: data,
	}
	buf, _ := proto.Marshal(&res)
	conn.codec.Encode(buf)
}
