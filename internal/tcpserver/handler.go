package tcpserver

import (
	"github.com/golang/protobuf/proto"
	"xim/internal/pb"
	tcppb "xim/internal/tcpserver/pb"
	"xim/internal/tcpserver/rpc"
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
	case tcppb.PacketType_PACKET_SEND_MESSAGE:
		h.Message(req)
	}
}

func (h *Handler) SignIn(req tcppb.RequestPacket) {
	var r pb.SignInReq
	proto.Unmarshal(req.Data, &r)

	ctx := h.ctx
	ctx.DeviceId = r.DeviceId
	ctx.UserId = r.UserId
	ctx.AppId = int8(r.App)

	err := rpc.SignIn(r)

	if err != nil {
		h.ResponseError(req, pb.ErrorCode_EC_INTERNAL_SERVICE, err.Error())
		return
	}

	// 保存连接
	manager.Store(r.DeviceId, ctx)
}

func (h *Handler) HeartBeat(req tcppb.RequestPacket) {
	h.ResponseOk(req)
}

func (h *Handler) Message(req tcppb.RequestPacket) {
	var msg pb.UpMessage
	proto.Unmarshal(req.Data, &msg)

	err := rpc.SendMessage(msg)
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
