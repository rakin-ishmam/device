package handler

import (
	"device/app/api/handler/device"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handlers represents the handlers
type Handlers struct {
	Device *device.Handler
}

// NewRouter creates a new router
func NewRouter(hs Handlers) http.Handler {
	r := chi.NewRouter()
	r.Route("/api/v1/devices", func(r chi.Router) {
		r.Get("/{id}", hs.Device.GetByID)
		r.Put("/{id}", hs.Device.Update)
		r.Delete("/{id}", hs.Device.Delete)
		r.Get("/", hs.Device.SearchByBrand)
		r.Post("/", hs.Device.Create)
	})
	return r
}
