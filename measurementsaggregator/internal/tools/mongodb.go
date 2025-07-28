package tools

import (
	"context"

	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client                *mongo.Client
	database              *mongo.Database
	deviceCollection      *mongo.Collection
	measurementCollection *mongo.Collection
}

// Connect implements DatabaseInterface.
func (m *mongoDB) Connect(ctx context.Context, url string) error {
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url /*"mongodb://foo:bar@localhost:27017"*/))

	if err != nil {
		return err
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	m.client = client

	log.Info("MongoClient connected")

	return nil
}

// SetupDatabase implements DatabaseInterface.
func (m *mongoDB) SetupDatabase(dbName string) error {
	m.database = m.client.Database(dbName)
	m.deviceCollection = m.database.Collection("devices")
	m.measurementCollection = m.database.Collection("measurements")

	return nil
}

func (m *mongoDB) GetStats(deviceId string, metric string) (*models.GeneralStats, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"deviceId": deviceId, "metric": metric}},
		{"$group": bson.M{
			"_id": bson.M{"deviceId": "$deviceId", "metric": "$metric"},
			"avg": bson.M{"$avg": "$value"},
			"min": bson.M{"$min": "$value"},
			"max": bson.M{"$max": "$value"},
		}},
	}

	cursor, err := m.measurementCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var stats *models.GeneralStats
	if cursor.Next(context.TODO()) {
		var result struct {
			ID struct {
				DeviceID string `bson:"deviceId"`
				Metric   string `bson:"metric"`
			} `bson:"_id"`
			Avg float64 `bson:"avg"`
			Min float64 `bson:"min"`
			Max float64 `bson:"max"`
		}

		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		stats = &models.GeneralStats{
			DeviceID: result.ID.DeviceID,
			Metric:   result.ID.Metric,
			Avg:      result.Avg,
			Min:      result.Min,
			Max:      result.Max,
		}
	}

	return stats, nil
}
