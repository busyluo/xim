package adapter

import (
	"context"
	"xim/internal/app/logic/entity"
)

type DeviceRepository interface {
	Register(ctx context.Context) error
	GetByUserID(ctx context.Context, id int64) (*entity.Device, error)
}
