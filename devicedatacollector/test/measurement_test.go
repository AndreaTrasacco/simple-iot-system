package test

import (
	"testing"
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/services"
)

func TestUploadMeasurement(t *testing.T) {
	mockDB := &MockDatabase{}
	services.InitMeasurementService(mockDB)

	measurement := &models.Measurement{
		DeviceID:  "dev-123",
		Timestamp: time.Now(),
		Metric:    "temperature",
		Value:     25.6,
	}

	err := services.UploadMeasurement(measurement)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(mockDB.Measurements) != 1 {
		t.Fatalf("Expected 1 measurement in mock DB, got %d", len(mockDB.Measurements))
	}
}

func TestGetMeasurements(t *testing.T) {
	mockDB := &MockDatabase{
		Measurements: []*models.Measurement{
			{
				DeviceID:  "dev-123",
				Timestamp: time.Now().Add(-1 * time.Hour),
				Metric:    "temperature",
				Value:     26.5,
			},
		},
	}
	services.InitMeasurementService(mockDB)

	fromTime := time.Now().Add(-2 * time.Hour)
	toTime := time.Now()

	measurements, err := services.GetMeasurements("dev-123", fromTime, toTime)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(measurements) != 1 {
		t.Fatalf("Expected 1 measurement, got %d", len(measurements))
	}

	if measurements[0].DeviceID != "dev-123" {
		t.Fatalf("Expected device ID 'dev-123', got '%s'", measurements[0].DeviceID)
	}
}
