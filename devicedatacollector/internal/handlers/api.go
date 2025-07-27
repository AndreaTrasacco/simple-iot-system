package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/devices", func(r chi.Router) {
		r.Post("/", RegisterDevice)
		r.Get("/", GetDevices)
	})

	r.Route("/measurements", func(r chi.Router) {
		r.Post("/", UploadMeasurements)
		r.Get("/", GetMeasurements)
	})
}
