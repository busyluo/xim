package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"xim/internal/logic/entity"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return DeviceRepository{
		db: db,
	}
}

func (DeviceRepository) Register(ctx context.Context) error {
	panic("implement me")
}

func (DeviceRepository) GetByUserID(ctx context.Context, id int64) (*entity.Device, error) {
	panic("implement me")
}
