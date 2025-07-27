package models

import (
	"time"
)

type Measurement struct {
	DeviceID  string    `json:"deviceId" bson:"deviceId"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Metric    string    `json:"metric" bson:"metric"`
	Value     float64   `json:"value" bson:"value"`
}
