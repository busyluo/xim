package tcpserver

import (
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net"
	"xim/internal/app/tcpserver/codec"
	"xim/internal/app/tcpserver/pb"
)

// ConnCtx context fo net.TCPConn
type ConnCtx struct {
	codec *codec.LvCodec

	UserId   int64
	DeviceId int64
	AppId    int8
}

func ServeConn(conn *net.TCPConn) {
	var ctx = &ConnCtx{
		codec: codec.NewLvCodec(conn),
	}

	for {
		data, err := ctx.codec.Decode()
		if err != nil {
			log.Error(err)
			return
		}
		var req pb.RequestPacket
		proto.Unmarshal(data, &req)

		handler.Handle(ctx, req)
	}
}
