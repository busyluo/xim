package service

import (
	"context"
	"xim/internal/app/logic/entity"
	"xim/internal/app/logic/repository1"
)

type DeviceService struct {
	repo repository1.Device
}

func NewDeviceService(repo repository1.Device) DeviceService {
	return DeviceService{repo: repo}
}

func (DeviceService) Register(ctx context.Context, device entity.Device) (int64, error) {
	panic("implement me")
}
