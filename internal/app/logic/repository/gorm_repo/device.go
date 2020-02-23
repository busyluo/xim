package gorm_repo

import (
	"context"
	"github.com/jinzhu/gorm"
	"xim/internal/app/logic/entity"
	"xim/internal/app/logic/repository/adapter"
)

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) adapter.DeviceRepository {
	return deviceRepository{
		db: db,
	}
}

func (r deviceRepository) NewDevice(ctx context.Context, d entity.Device) error {
	r.db.Create(&d)
	return nil
}

func (deviceRepository) GetByUserID(ctx context.Context, id int64) (*entity.Device, error) {
	panic("implement me")
}
