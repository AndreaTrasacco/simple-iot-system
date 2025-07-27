package tools

import (
	"context"
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
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
	m.deviceCollection = m.database.Collection("devices")           // Better if "devices" is taken from a configuration file
	m.measurementCollection = m.database.Collection("measurements") // Better if "measurements" is taken from a configuration file

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "deviceId", Value: 1}}, // Ascending index
		Options: options.Index().SetUnique(true),     // Ensure deviceId is unique
	}

	// Create unique index on deviceId
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	_, err = m.deviceCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Set a compound index model
	indexModel = mongo.IndexModel{
		Keys:    bson.D{{Key: "deviceId", Value: 1}, {Key: "timestamp", Value: 1}}, // Ascending index for both fields
		Options: options.Index(),                                                   // Add more options if needed
	}

	// Create compound index on deviceId and timestamp
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = m.measurementCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// SaveDevice implements DatabaseInterface.
func (m *mongoDB) SaveDevice(device *models.Device) error {
	_, err := m.deviceCollection.InsertOne(context.TODO(), device)
	return err
}

// SaveMeasurement implements DatabaseInterface.
func (m *mongoDB) SaveMeasurement(measurement *models.Measurement) error {
	_, err := m.measurementCollection.InsertOne(context.TODO(), measurement)
	return err
}

// GetDevices implements DatabaseInterface.
func (m *mongoDB) GetDevices() ([]*models.Device, error) {
	var devices []*models.Device
	cursor, err := m.deviceCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &devices); err != nil {
		return nil, err
	}
	return devices, nil
}

// GetMeasurements implements DatabaseInterface.
func (m *mongoDB) GetMeasurementsByDeviceAndTimestampRange(deviceId string, from, to time.Time) ([]*models.Measurement, error) {
	var measurements []*models.Measurement
	filter := bson.M{
		"deviceId": deviceId,
		"timestamp": bson.M{
			"$gte": from,
			"$lte": to,
		},
	}
	cursor, err := m.measurementCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &measurements); err != nil {
		return nil, err
	}
	return measurements, nil
}
