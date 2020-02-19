package service

import (
	"context"
	"github.com/google/wire"
	"xim/internal/logic/entity"
	"xim/internal/pb"
)

type DeviceInterface interface {
	Register(ctx context.Context, device entity.Device) (int64, error)
}

type MessageInterface interface {
	Register(ctx context.Context, device entity.Device) (int64, error)
	SendMessage(ctx context.Context, message *pb.UpMessage) error
}

var DeviceServiceProvider = wire.NewSet(NewDeviceService, wire.Bind(new(DeviceInterface), new(DeviceService)))
