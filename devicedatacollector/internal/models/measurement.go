package models

import (
	"time"
)

type Measurement struct {
	DeviceID  string    `json:"deviceId" bson:"deviceId" validate:"required"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp" validate:"required"`
	Metric    string    `json:"metric" bson:"metric" validate:"required"`
	Value     float64   `json:"value" bson:"value" validate:"required"`
}
