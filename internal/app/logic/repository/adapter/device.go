package adapter

import (
	"context"
	"xim/internal/app/logic/entity"
)

type DeviceRepository interface {
	NewDevice(ctx context.Context, device entity.Device) error
	GetByUserID(ctx context.Context, id int64) (*entity.Device, error)
}
