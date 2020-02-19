package repository

import (
	"context"
	"github.com/google/wire"
	"xim/internal/logic/entity"
	"xim/internal/logic/repository/mysql"
)

type Device interface {
	Register(ctx context.Context) error
	GetByUserID(ctx context.Context, id int64) (*entity.Device, error)
}

var DeviceRepositoryProvider = wire.NewSet(mysql.NewDeviceRepository, wire.Bind(new(Device), new(mysql.DeviceRepository)))
