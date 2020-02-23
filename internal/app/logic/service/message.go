package service

import (
	"context"
	"xim/internal/app/logic/cache"
	"xim/internal/app/logic/entity"
	"xim/internal/pkg/local_call"
	"xim/internal/pkg/pb"
)

type MessageImpl struct{}

var MessageService = new(MessageImpl)

func (MessageImpl) Register(ctx context.Context, device entity.Device) (int64, error) {
	panic("implement me")
}

// 消息中转
func (MessageImpl) SendMessage(ctx context.Context, up *pb.UpMessage) error {
	devices, err := cache.GetAllOnlineDevice(up.Receiver)
	if err != nil {
		return err
	}

	user := ctx.Value("user").(int64)

	var down pb.DownMessage
	down.Sender = user
	down.Content = up.Content
	down.SendTime = up.SendTime
	down.Seq = up.Seq

	req := pb.DispatchReq{
		Id:  devices,
		Msg: &down,
	}
	local_call.TcpServer.DispatchMessage(ctx, &req)
	return nil
}
