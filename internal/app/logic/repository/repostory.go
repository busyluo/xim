package repository

import (
	"github.com/jinzhu/gorm"
	"xim/internal/app/logic/repository/adapter"
	"xim/internal/app/logic/repository/gorm_repo"
)

var Device adapter.DeviceRepository

func InitOrmRepository(db *gorm.DB) {
	Device = gorm_repo.NewDeviceRepository(db)
}

// if mock
func InitMockRepository() {
	// Device = gorm_repo.NewDeviceRepository(db)
}
