package services

import (
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/tools"
)

func InitMeasurementService(dbRef tools.DatabaseInterface) {
	db = dbRef
}

func UploadMeasurement(measurement *models.Measurement) error {
	return db.SaveMeasurement(measurement)
}

func GetMeasurements(deviceId string, from, to time.Time) ([]*models.Measurement, error) {
	return db.GetMeasurementsByDeviceAndTimestampRange(deviceId, from, to)
}
