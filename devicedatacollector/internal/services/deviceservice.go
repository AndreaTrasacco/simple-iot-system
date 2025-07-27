package services

import (
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/tools"
)

var db tools.DatabaseInterface

func InitDeviceService(dbRef tools.DatabaseInterface) {
	db = dbRef
}

func CreateDevice(device *models.Device) error {
	device.CreatedAt = time.Now()
	return db.SaveDevice(device)
}

func GetAllDevices() ([]*models.Device, error) {
	return db.GetDevices()
}
