package service

import (
	"context"
	"xim/internal/logic/entity"
	"xim/internal/pb"
)

type MessageImpl struct{}

var MessageService = new(MessageImpl)

func (MessageImpl) Register(ctx context.Context, device entity.Device) (int64, error) {
	panic("implement me")
}

func (MessageImpl) SendMessage(ctx context.Context, message *pb.UpMessage) error {
	// get device id by user id.
}
