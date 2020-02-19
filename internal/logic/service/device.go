package service

import (
	"context"
	"xim/internal/logic/entity"
	"xim/internal/logic/repository"
)

type DeviceService struct {
	repo repository.Device
}

func NewDeviceService(repo repository.Device) DeviceService {
	return DeviceService{repo: repo}
}

func (DeviceService) Register(ctx context.Context, device entity.Device) (int64, error) {
	panic("implement me")
}
