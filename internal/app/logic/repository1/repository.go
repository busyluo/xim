package repository1

import (
	"context"
	"github.com/google/wire"
	"xim/internal/app/logic/entity"
	"xim/internal/app/logic/repository/mysql"
)

type Device interface {
	Register(ctx context.Context) error
	GetByUserID(ctx context.Context, id int64) (*entity.Device, error)
}

var DeviceRepositoryProvider = wire.NewSet(mysql.NewDeviceRepository, wire.Bind(new(Device), new(mysql.DeviceRepository)))
