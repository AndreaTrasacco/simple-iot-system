package tools

import (
	"context"
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	log "github.com/sirupsen/logrus"
)

type DatabaseInterface interface {
	Connect(ctx context.Context, url string) error
	SetupDatabase(dbName string) error
	SaveDevice(dev *models.Device) error
	GetDevices() ([]*models.Device, error)
	SaveMeasurement(measurement *models.Measurement) error
	GetMeasurementsByDeviceAndTimestampRange(deviceId string, from, to time.Time) ([]*models.Measurement, error)
}

func NewDatabase(ctx context.Context, url string) (*DatabaseInterface, error) {

	var database DatabaseInterface = &mongoDB{}
	var err error = database.Connect(ctx, url)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = database.SetupDatabase("a")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
