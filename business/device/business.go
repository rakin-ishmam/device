package device

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

// Store is an interface to interact with the database
type Store interface {
	ByID(ctx context.Context, id string) (Device, error)
	Create(ctx context.Context, d Device) error
	Update(ctx context.Context, id string, data UpdateDevice) error
	GetAll(ctx context.Context, offset, limit int) ([]Device, error)
	Delete(ctx context.Context, id string) error
	SearchByBrand(ctx context.Context, brand string, offset, limit int) ([]Device, error)
}

// Business is the business logic for the device
type Business struct {
	store Store
}

// NewBusiness creates a new business logic for the device
func NewBusiness(store Store) *Business {
	return &Business{
		store: store,
	}
}

// GetByID returns a device by its ID
func (b *Business) Create(ctx context.Context, cd CreateDevice) (Device, error) {
	d := cd.toDevice()
	if err := b.store.Create(ctx, d); err != nil {
		return Device{}, fmt.Errorf("store.Create: %w", err)
	}
	return d, nil
}

// Update updates a device
func (b *Business) Update(ctx context.Context, id string, data UpdateDevice) error {
	if err := b.store.Update(ctx, id, data); err != nil {
		return fmt.Errorf("store.Update: %w", err)
	}
	return nil
}

func (b *Business) GetByID(ctx context.Context, id string) (Device, error) {
	device, err := b.store.ByID(ctx, id)
	if err != nil {
		return Device{}, fmt.Errorf("store.ByID: %w", err)
	}
	return device, nil
}

// Getall returns all devices
func (b *Business) GetAll(ctx context.Context, offset, limit int) ([]Device, error) {
	devices, err := b.store.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("store.GetAll: %w", err)
	}
	return devices, nil
}

// Delete deletes a device
func (b *Business) Delete(ctx context.Context, id string) error {
	if err := b.store.Delete(ctx, id); err != nil {
		return fmt.Errorf("store.Delete: %w", err)
	}
	return nil
}

// SearchByBrand returns all devices by a brand
func (b *Business) SearchByBrand(ctx context.Context, brand string, offset, limit int) ([]Device, error) {
	devices, err := b.store.SearchByBrand(ctx, brand, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("store.SearchByBrand: %w", err)
	}
	return devices, nil
}
