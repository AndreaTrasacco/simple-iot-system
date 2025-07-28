package models

import (
	"time"
)

type Device struct {
	DeviceID  string    `json:"deviceId" bson:"deviceId" validate:"required"`
	Type      string    `json:"type" bson:"type" validate:"required"`
	Location  string    `json:"location" bson:"location" validate:"required"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
