package service

import (
	"github.com/google/wire"
	"xim/internal/logic/repository"
)

func CreateDeviceService(repo repository.Device) (DeviceService, error) {
	panic(wire.Build(DeviceServiceProvider))
}
