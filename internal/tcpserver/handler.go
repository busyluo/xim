package tcpserver

import (
	"github.com/golang/protobuf/proto"
	"xim/internal/tcpserver/pb"
	"xim/internal/tcpserver/rpc"
)

type Handler struct {
	ctx *ConnCtx
}

var handler Handler

func (h *Handler) Handle(ctx *ConnCtx, req pb.RequestPacket) {
	h.ctx = ctx

	switch req.Type {
	case pb.PacketType_PACKET_LOGIN:
		h.Login(req)
	case pb.PacketType_PACKET_HEARTBEAT:
		h.HeartBeat(req)
	case pb.PacketType_PACKET_SEND_MESSAGE:
		h.Message(req)
	}
}

func (h *Handler) Login(req pb.RequestPacket) {
	var login pb.LoginReq
	proto.Unmarshal(req.Data, &login)

	ctx := h.ctx
	ctx.DeviceId = login.DeviceId
	ctx.UserId = login.UserId
	ctx.AppId = int8(login.App)

	err := rpc.Login(login)

	if err != nil {
		h.ResponseError(req, pb.ErrorCode_EC_INTERNAL_SERVICE, err.Error())
		return
	}

	// 保存连接
	manager.Store(login.DeviceId, ctx)
}

func (h *Handler) HeartBeat(req pb.RequestPacket) {
	h.ResponseOk(req)
}

func (h *Handler) Message(req pb.RequestPacket) {
	var msg pb.SendMessage
	proto.Unmarshal(req.Data, &msg)

	err := rpc.SendMessage(msg)
	if err != nil {
		h.ResponseError(req, pb.ErrorCode_EC_INTERNAL_SERVICE, err.Error())
		return
	}

	h.ResponseOk(req)
}

func (h *Handler) ResponseOk(req pb.RequestPacket) {
	res := pb.ResponsePacket{
		Type:    req.Type,
		ReqId:   req.Id,
		Code:    pb.ErrorCode_EC_SUCCESS,
		Message: "",
		Data:    nil,
	}
	buf, _ := proto.Marshal(&res)
	h.ctx.codec.Encode(buf)
}

func (h *Handler) ResponseError(req pb.RequestPacket, ec pb.ErrorCode, msg string) {
	res := pb.ResponsePacket{
		Type:    req.Type,
		ReqId:   req.Id,
		Code:    ec,
		Message: msg,
		Data:    nil,
	}
	buf, _ := proto.Marshal(&res)
	h.ctx.codec.Encode(buf)
}
