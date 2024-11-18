package device

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Device represents a device
type Device struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateDevice represents the data needed to create a device
type CreateDevice struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

// Validate validates the CreateDevice fields
func (cd CreateDevice) Validate() error {
	if cd.Name == "" {
		return fmt.Errorf("name is required")
	}
	if cd.Brand == "" {
		return fmt.Errorf("brand is required")
	}
	return nil
}

// toDevice converts a CreateDevice to a Device
func (cd CreateDevice) toDevice() Device {
	return Device{
		ID:        uuid.NewString(),
		Name:      cd.Name,
		Brand:     cd.Brand,
		CreatedAt: time.Now(),
	}
}

// UpdateDevice represents the data needed to update a device
type UpdateDevice struct {
	Name  *string `json:"name,omitempty"`
	Brand *string `json:"brand,omitempty"`
}
