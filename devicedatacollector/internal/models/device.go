package models

import (
	"time"
)

type Device struct {
	DeviceID  string    `json:"deviceId" bson:"deviceId"`
	Type      string    `json:"type" bson:"type"`
	Location  string    `json:"location" bson:"location"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
