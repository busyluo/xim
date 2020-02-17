package main

import "github.com/labstack/echo/v4"

// DeviceRequest 设备信息
type DeviceRequest struct {
	// 系统
	Os string `json:"os"`
	// 设备名
	Name string `json:"name"`
	// 厂商
	Brand string `json:"brand"`
	// 标识(手机MEID或电脑MAC)
	Identifier string `json:"identifier"`
}

func device(c echo.Context) error {
	var req DeviceRequest
	err := c.Bind(req)
	if err != nil {
		return err

	}
	return nil
}
