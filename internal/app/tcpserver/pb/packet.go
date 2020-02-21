package pb

import (
	"github.com/golang/protobuf/proto"
	"xim/internal/pkg/pb"
)

type RequestId int32

func MarshalRequest(id RequestId, message proto.Message) []byte {
	var t PacketType
	switch message.(type) {
	case *pb.SignInReq:
		t = PacketType_PACKET_SIGN_IN
	default:
		panic("unexpected message")
	}
	data, _ := proto.Marshal(message)
	r := &RequestPacket{
		Type: t,
		Id:   int32(id),
		Data: data,
	}
	buf, _ := proto.Marshal(r)
	return buf
}
