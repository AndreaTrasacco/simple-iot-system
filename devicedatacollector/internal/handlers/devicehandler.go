package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/api"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/models"
	"github.com/AndreaTrasacco/simple-iot-system/devicedatacollector/internal/services"
	log "github.com/sirupsen/logrus"
)

func RegisterDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	if err := services.CreateDevice(&device); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := services.GetAllDevices()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	if len(devices) == 0 {
		json.NewEncoder(w).Encode(map[string]string{})
		return
	}

	json.NewEncoder(w).Encode(devices)
}
