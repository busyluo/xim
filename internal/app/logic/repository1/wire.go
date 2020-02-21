package repository1

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func CreateDeviceRepository(db *gorm.DB) (Device, error) {
	panic(wire.Build(DeviceRepositoryProvider))
}
