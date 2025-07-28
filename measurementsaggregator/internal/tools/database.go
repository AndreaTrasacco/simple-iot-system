package tools

import (
	"context"

	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/models"
	log "github.com/sirupsen/logrus"
)

type DatabaseInterface interface {
	Connect(ctx context.Context, url string) error
	SetupDatabase(dbName string) error
	GetStats(deviceId string, metric string) (*models.GeneralStats, error)
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
