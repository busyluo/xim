package tcpserver

import (
	"github.com/golang/protobuf/proto"
	"xim/internal/tcpserver/pb"
)

type Handler struct{}

var handler Handler

func (h *Handler) Handle(ctx *ConnCtx, req pb.RequestPacket) {
	switch req.Type {
	case pb.PacketType_LOGIN:
		h.Login(ctx, req)
	case pb.PacketType_HEARTBEAT:
		h.HeartBeat(ctx, req)
	}
}

func (h *Handler) Login(ctx *ConnCtx, req pb.RequestPacket) {
	var login pb.LoginReq
	proto.Unmarshal(req.Data, &login)
	manager.Store(login.DeviceId, ctx)
}

func (h *Handler) HeartBeat(ctx *ConnCtx, req pb.RequestPacket) {
	//var login pb.LoginReq
	//proto.Unmarshal(req.Data, &login)
	//manager.Store(login.DeviceId, ctx)
}
