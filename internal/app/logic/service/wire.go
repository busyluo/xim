package service

import (
	"github.com/google/wire"
	"xim/internal/app/logic/repository"
)

func CreateDeviceService(repo repository1.Device) (DeviceService, error) {
	panic(wire.Build(DeviceServiceProvider))
}
