package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/api"
	"github.com/AndreaTrasacco/simple-iot-system/measurementsaggregator/internal/services"
	log "github.com/sirupsen/logrus"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	deviceId := r.URL.Query().Get("deviceId")
	metric := r.URL.Query().Get("metric")

	if deviceId == "" || metric == "" {
		log.Error("Wrong parameters values: deviceId=" + deviceId + "; metric=" + metric + ";")
		api.RequestErrorHandler(w, fmt.Errorf("deviceId and metric query parameters are required"))
		return
	}

	stats, err := services.GetStats(deviceId, metric)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	if stats == nil {
		json.NewEncoder(w).Encode(map[string]string{})
		return
	}

	json.NewEncoder(w).Encode(stats)
}
