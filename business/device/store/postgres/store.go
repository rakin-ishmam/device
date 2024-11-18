package postgres

import (
	"context"
	"device/business/device"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Store is a postgres implementation of the device.Store
type Store struct {
	db *gorm.DB
}

// this is a compile time check to ensure Store implements device.Store
var _ device.Store = (*Store)(nil)

// NewStore creates a new Store instance
func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// ByID returns a device by its ID
func (s *Store) ByID(ctx context.Context, id string) (device.Device, error) {
	var d Device
	result := s.db.WithContext(ctx).First(&d, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return device.Device{}, device.ErrNotFound
	}
	if result.Error != nil {
		return device.Device{}, fmt.Errorf("db.First[%s]: %w", id, result.Error)
	}

	return toBusinessDevice(d), nil
}

// Create creates a new device
func (s *Store) Create(ctx context.Context, d device.Device) error {
	createData := fromBusinessDevice(d)
	result := s.db.WithContext(ctx).Create(&createData)
	if result.Error != nil {
		return fmt.Errorf("db.Create: %w", result.Error)
	}
	return nil
}

// Update updates a device
func (s *Store) Update(ctx context.Context, id string, data device.UpdateDevice) error {
	updates := map[string]interface{}{}
	if data.Name != nil {
		updates["name"] = *data.Name
	}
	if data.Brand != nil {
		updates["brand"] = *data.Brand
	}

	result := s.db.WithContext(ctx).Model(&Device{}).Where("id = ?", id).Updates(updates)
	if result.RowsAffected == 0 {
		return device.ErrNotFound
	}
	if result.Error != nil {
		return fmt.Errorf("db.Updates[%s]: %w", id, result.Error)
	}
	return nil
}

// GetAll returns all devices
func (s *Store) GetAll(ctx context.Context, offset, limit int) ([]device.Device, error) {
	var devices []Device
	result := s.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&devices)
	if result.Error != nil {
		return nil, fmt.Errorf("db.Find: %w", result.Error)
	}
	return toBusinessDevices(devices), result.Error
}

// Delete deletes a device
func (s *Store) Delete(ctx context.Context, id string) error {
	result := s.db.WithContext(ctx).Delete(&Device{ID: id}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("db.Delete[%s]: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return device.ErrNotFound
	}
	return nil
}

// SearchByBrand searches devices by brand
func (s *Store) SearchByBrand(ctx context.Context, brand string, offset, limit int) ([]device.Device, error) {
	var devices []Device
	result := s.db.WithContext(ctx).Where("brand = ?", brand).Offset(offset).Limit(limit).Find(&devices)
	if result.Error != nil {
		return nil, fmt.Errorf("db.Find: %w", result.Error)
	}
	return toBusinessDevices(devices), nil
}
