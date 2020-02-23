package service

import (
	"context"
	"xim/internal/app/logic/entity"
	"xim/internal/app/logic/repository"
)

type device struct{}

var Device device

func (device) NewDevice(ctx context.Context, d entity.Device) error {
	return repository.Device.NewDevice(ctx, d)
}

func (device) Register(ctx context.Context, device entity.Device) (int64, error) {
	panic("implement me")
}
