package logic

import (
	"context"
	"github.com/golang/protobuf/proto"
	tcppb "xim/internal/app/tcpserver/pb"
	"xim/internal/app/tcpserver/rpc"
	"xim/internal/pkg/pb"
)

type Handler struct {
	ctx *ConnCtx
}

var handler Handler

func (h *Handler) Handle(ctx *ConnCtx, req tcppb.RequestPacket) {
	h.ctx = ctx

	switch req.Type {
	case tcppb.PacketType_PACKET_SIGN_IN:
		h.SignIn(req)
	case tcppb.PacketType_PACKET_HEARTBEAT:
		h.HeartBeat(req)
	case tcppb.PacketType_PACKET_MESSAGE:
		h.Message(req)
	}
}

func (h *Handler) SignIn(req tcppb.RequestPacket) {
	var r pb.SignInReq
	proto.Unmarshal(req.Data, &r)

	conn := h.ctx
	conn.DeviceId = r.DeviceId
	conn.UserId = r.UserId
	conn.AppId = int8(r.App)

	ctx := context.WithValue(context.Background(), "user", h.ctx.UserId)
	err := rpc.SignIn(ctx, r)

	if err != nil {
		h.ResponseError(req, pb.ErrorCode_EC_INTERNAL_SERVICE, err.Error())
		return
	}

	// 保存连接
	manager.Store(r.DeviceId, conn)
}

func (h *Handler) HeartBeat(req tcppb.RequestPacket) {
	h.ResponseOk(req)
}

func (h *Handler) Message(req tcppb.RequestPacket) {
	var msg pb.UpMessage
	proto.Unmarshal(req.Data, &msg)

	ctx := context.WithValue(context.Background(), "user", h.ctx.UserId)
	err := rpc.SendMessage(ctx, msg)
	if err != nil {
		h.ResponseError(req, pb.ErrorCode_EC_INTERNAL_SERVICE, err.Error())
		return
	}

	h.ResponseOk(req)
}

func (h *Handler) ResponseOk(req tcppb.RequestPacket) {
	res := tcppb.ResponsePacket{
		Type:    req.Type,
		ReqId:   req.Id,
		Code:    int32(pb.ErrorCode_EC_SUCCESS),
		Message: "",
		Data:    nil,
	}
	buf, _ := proto.Marshal(&res)
	h.ctx.codec.Encode(buf)
}

func (h *Handler) ResponseError(req tcppb.RequestPacket, ec pb.ErrorCode, msg string) {
	res := tcppb.ResponsePacket{
		Type:    req.Type,
		ReqId:   req.Id,
		Code:    int32(ec),
		Message: msg,
		Data:    nil,
	}
	buf, _ := proto.Marshal(&res)
	h.ctx.codec.Encode(buf)
}
