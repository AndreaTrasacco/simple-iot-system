package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/handlers"
	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/services"
	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/tools"
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

	services.InitStatsService(*dbInstance)

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting data aggregator API service...")

	var aggregatorPort = os.Getenv("AGGREGATOR_PORT")
	if aggregatorPort == "" {
		log.Error("AGGREGATOR_PORT environment variable undefined")
		return
	}
	errHttp := http.ListenAndServe(":"+aggregatorPort, r)
	if errHttp != nil {
		log.Error(errHttp)
	}
}
