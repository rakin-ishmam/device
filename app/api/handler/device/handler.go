package device

import (
	"context"
	"device/business/device"
	"device/pkg/web"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Business represents the device business interface
type Business interface {
	Create(ctx context.Context, cd device.CreateDevice) (device.Device, error)
	Update(ctx context.Context, id string, data device.UpdateDevice) error
	GetByID(ctx context.Context, id string) (device.Device, error)
	GetAll(ctx context.Context, offset, limit int) ([]device.Device, error)
	Delete(ctx context.Context, id string) error
	SearchByBrand(ctx context.Context, brand string, offset, limit int) ([]device.Device, error)
}

// Handler represents the device handler
type Handler struct {
	business Business
}

// NewHandler creates a new device handler
func NewHandler(b Business) *Handler {
	return &Handler{
		business: b,
	}
}

// Create creates a device
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var createDevice device.CreateDevice
	if err := json.NewDecoder(r.Body).Decode(&createDevice); err != nil {
		web.SendError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload"))
		return
	}

	if err := createDevice.Validate(); err != nil {
		web.SendError(w, http.StatusBadRequest, err)
		return
	}

	device, err := h.business.Create(r.Context(), createDevice)
	if err != nil {
		web.SendError(w, http.StatusInternalServerError, fmt.Errorf("unable to create device"))
		logrus.WithError(fmt.Errorf("business.Create: %w", err)).Error("unable to create device")
		return
	}

	web.SendCreated(w, device)
}

// Update updates a device
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		web.SendError(w, http.StatusBadRequest, err)
		return
	}

	var updateDevice device.UpdateDevice
	if err := json.NewDecoder(r.Body).Decode(&updateDevice); err != nil {
		web.SendError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload"))
		return
	}

	err = h.business.Update(r.Context(), id, updateDevice)
	if err == nil {
		web.SendOk(w, map[string]string{
			"message": "device updated",
		})
		return
	}

	if errors.Is(err, device.ErrNotFound) {
		web.SendError(w, http.StatusNotFound, fmt.Errorf("device not found"))
		return
	}
	web.SendError(w, http.StatusInternalServerError, fmt.Errorf("unable to update device"))
	logrus.WithError(fmt.Errorf("business.Update: %w", err)).Error("unable to update device")
}

// GetByID returns a device by its ID
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		web.SendError(w, http.StatusBadRequest, err)
		return
	}

	d, err := h.business.GetByID(r.Context(), id)
	if err == nil {
		web.SendOk(w, d)
		return
	}

	if errors.Is(err, device.ErrNotFound) {
		web.SendError(w, http.StatusNotFound, fmt.Errorf("device not found"))
		return
	}

	web.SendError(w, http.StatusInternalServerError, fmt.Errorf("unable to get device"))
	logrus.WithError(fmt.Errorf("business.GetByID: %w", err)).Error("unable to get device")
}

// Delete deletes a device
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		web.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = h.business.Delete(r.Context(), id)
	if err == nil {
		web.SendOk(w, map[string]string{
			"message": "device deleted",
		})
		return
	}

	if errors.Is(err, device.ErrNotFound) {
		web.SendError(w, http.StatusNotFound, fmt.Errorf("device not found"))
		return
	}

	web.SendError(w, http.StatusInternalServerError, fmt.Errorf("unable to delete device"))
	logrus.WithError(fmt.Errorf("business.Delete: %w", err)).Error("unable to delete device")
}

// GetAll returns all devices
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	offset, limit := web.ParsePaginationParams(r)

	devices, err := h.business.GetAll(r.Context(), offset, limit)
	if err != nil {
		web.SendError(w, http.StatusInternalServerError, fmt.Errorf("unable to get devices"))
		logrus.WithError(fmt.Errorf("business.GetAll: %w", err)).Error("unable to get devices")
		return
	}

	web.SendOk(w, devices)
}

// SearchByBrand searches devices by brand
func (h *Handler) SearchByBrand(w http.ResponseWriter, r *http.Request) {
	brand := web.ParseStrQuery("brand", r)
	if brand == "" {
		h.GetAll(w, r)
		return
	}

	offset, limit := web.ParsePaginationParams(r)

	devices, err := h.business.SearchByBrand(r.Context(), brand, offset, limit)
	if err != nil {
		web.SendError(w, http.StatusInternalServerError, fmt.Errorf("unable to search devices by brand"))
		logrus.WithError(fmt.Errorf("business.SearchByBrand: %w", err)).Error("unable to search devices by brand")
		return
	}

	web.SendOk(w, devices)
}

// parseID parses the ID from the request
func parseID(r *http.Request) (string, error) {
	id := web.ParseStrURLParam("id", r)
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	_, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("id is not a valid UUID")
	}
	return id, nil
}
