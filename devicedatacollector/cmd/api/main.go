package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/handlers"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/services"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	var dbUrl = os.Getenv("MONGODB_URL")
	if dbUrl == "" {
		log.Error("MONGODB_URL environment variable undefined")
		return
	}
	dbInstance, err := tools.NewDatabase(ctx, dbUrl)
	if err != nil {
		log.Error(err)
		return
	}

	services.InitDeviceService(*dbInstance)
	services.InitMeasurementService(*dbInstance)

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting data collector API service...")

	var collectorPort = os.Getenv("COLLECTOR_PORT")
	if collectorPort == "" {
		log.Error("COLLECTOR_PORT environment variable undefined")
		return
	}
	errHttp := http.ListenAndServe(":"+collectorPort, r)
	if errHttp != nil {
		log.Error(errHttp)
	}
}
