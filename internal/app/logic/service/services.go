package service

import (
	"context"
	"xim/internal/app/logic/entity"
	"xim/internal/pkg/pb"
)

type DeviceInterface interface {
	NewDevice(ctx context.Context, device entity.Device) (int64, error)
	Register(ctx context.Context, device entity.Device) (int64, error)
}

type MessageInterface interface {
	Register(ctx context.Context, device entity.Device) (int64, error)
	SendMessage(ctx context.Context, message *pb.UpMessage) error
}

type UserInterface interface {
	NewUser(ctx context.Context, user entity.User) error
	Login(ctx context.Context)
}

// var DeviceServiceProvider = wire.NewSet(NewDeviceService, wire.Bind(new(DeviceInterface), new(DeviceService)))
