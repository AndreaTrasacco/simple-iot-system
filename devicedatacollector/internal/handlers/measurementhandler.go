package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/api"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/services"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/tools/validation"
	log "github.com/sirupsen/logrus"
)

func UploadMeasurements(w http.ResponseWriter, r *http.Request) {
	var measurements []models.Measurement
	if err := json.NewDecoder(r.Body).Decode(&measurements); err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	for _, measurement := range measurements {
		if err := validation.ValidateStruct(measurement); err != nil {
			log.Error(err)
			api.RequestErrorHandler(w, fmt.Errorf("malformed request! check body"))
			return
		}
	}

	for _, measurement := range measurements {
		if err := services.UploadMeasurement(&measurement); err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func GetMeasurements(w http.ResponseWriter, r *http.Request) {
	deviceId := r.URL.Query().Get("deviceId")
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	if deviceId == "" || fromStr == "" || toStr == "" {
		log.Error("Wrong parameters values: deviceId=" + deviceId + "; from=" + fromStr + "; to=" + toStr + ";")
		api.RequestErrorHandler(w, fmt.Errorf("deviceId, from, and to parameters are required"))
		return
	}

	fromTime, err := time.Parse(time.RFC3339, fromStr)
	if err != nil {
		log.Error("Invalid 'from' date format.")
		api.RequestErrorHandler(w, fmt.Errorf("invalid 'from' date format. Use RFC3339"))
		return
	}

	toTime, err := time.Parse(time.RFC3339, toStr)
	if err != nil {
		log.Error("Invalid 'to' date format.")
		api.RequestErrorHandler(w, fmt.Errorf("invalid 'to' date format. Use RFC3339"))
		return
	}

	measurements, err := services.GetMeasurements(deviceId, fromTime, toTime)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	if len(measurements) == 0 {
		json.NewEncoder(w).Encode(map[string]string{})
		return
	}

	json.NewEncoder(w).Encode(measurements)
}
