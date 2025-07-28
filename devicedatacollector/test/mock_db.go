package test

import (
	"context"
	"errors"
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
)

type MockDatabase struct {
	Devices      []*models.Device
	Measurements []*models.Measurement
	ShouldFail   bool
}

func (m *MockDatabase) Connect(ctx context.Context, url string) error {
	if m.ShouldFail {
		return errors.New("failed to connect")
	}
	return nil
}

func (m *MockDatabase) SetupDatabase(dbName string) error {
	if m.ShouldFail {
		return errors.New("failed to setup database")
	}
	return nil
}

func (m *MockDatabase) SaveDevice(dev *models.Device) error {
	if m.ShouldFail {
		return errors.New("failed to save device")
	}
	m.Devices = append(m.Devices, dev)
	return nil
}

func (m *MockDatabase) GetDevices() ([]*models.Device, error) {
	if m.ShouldFail {
		return nil, errors.New("failed to get devices")
	}
	return m.Devices, nil
}

func (m *MockDatabase) SaveMeasurement(measurement *models.Measurement) error {
	if m.ShouldFail {
		return errors.New("failed to save measurement")
	}
	m.Measurements = append(m.Measurements, measurement)
	return nil
}

func (m *MockDatabase) GetMeasurementsByDeviceAndTimestampRange(deviceId string, from, to time.Time) ([]*models.Measurement, error) {
	if m.ShouldFail {
		return nil, errors.New("failed to get measurements")
	}

	var filtered [](*models.Measurement)
	for _, measurement := range m.Measurements {
		if measurement.DeviceID == deviceId && measurement.Timestamp.After(from) && measurement.Timestamp.Before(to) {
			filtered = append(filtered, measurement)
		}
	}
	return filtered, nil
}
