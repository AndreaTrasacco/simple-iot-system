package test

import (
	"testing"
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/services"
)

func TestCreateDevice(t *testing.T) {
	mockDB := &MockDatabase{}
	services.InitDeviceService(mockDB)

	device := &models.Device{DeviceID: "dev-123", Type: "thermometer", Location: "43.717992, 10.946594"}

	err := services.CreateDevice(device)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(mockDB.Devices) != 1 {
		t.Fatalf("Expected 1 device in mock DB, got %d", len(mockDB.Devices))
	}
}

func TestGetDevices(t *testing.T) {
	mockDB := &MockDatabase{
		Devices: []*models.Device{
			{DeviceID: "dev-123", Type: "thermometer", Location: "43.717992, 10.946594", CreatedAt: time.Now()},
		},
	}
	services.InitDeviceService(mockDB)

	devices, err := services.GetAllDevices()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(devices) != 1 {
		t.Fatalf("Expected 1 device, got %d", len(devices))
	}

	if devices[0].DeviceID != "dev-123" {
		t.Fatalf("Expected device ID 'dev-123', got '%s'", devices[0].DeviceID)
	}
}

func TestCreateDeviceFailure(t *testing.T) {
	mockDB := &MockDatabase{ShouldFail: true} // Simulate failure
	services.InitDeviceService(mockDB)

	device := &models.Device{DeviceID: "dev-123", Type: "thermometer", Location: "43.717992, 10.946594"}

	err := services.CreateDevice(device)

	if err == nil {
		t.Fatal("Expected an error, but got none")
	}
}
