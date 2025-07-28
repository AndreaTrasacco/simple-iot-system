package services

import (
	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/tools"
)

var db tools.DatabaseInterface

func InitStatsService(dbRef tools.DatabaseInterface) {
	db = dbRef
}

func GetStats(deviceId string, metric string) (*models.GeneralStats, error) {
	return db.GetStats(deviceId, metric)
}
