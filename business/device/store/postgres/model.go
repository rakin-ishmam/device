package postgres

import (
	"device/business/device"
	"time"
)

// device represents a device
type Device struct {
	ID        string    `gorm:"primaryKey;column:id"` // Primary Key with custom column name
	Name      string    `gorm:"column:name"`          // Custom column name
	Brand     string    `gorm:"column:brand"`         // Custom column name
	CreatedAt time.Time `gorm:"column:created_at"`    // Explicit column name
}

// TableName overrides the default table name (from "devices" to "custom_devices").
func (Device) TableName() string {
	return "devices"
}

// toBusinessDevice converts a Device to a device.Device
func toBusinessDevice(d Device) device.Device {
	return device.Device{
		ID:        d.ID,
		Name:      d.Name,
		Brand:     d.Brand,
		CreatedAt: d.CreatedAt,
	}
}

// toBusinessDevices converts a slice of Device to a slice of device.Device
func toBusinessDevices(ds []Device) []device.Device {
	devices := make([]device.Device, len(ds))
	for i, d := range ds {
		devices[i] = toBusinessDevice(d)
	}
	return devices
}

// fromBusinessDevice converts a device.Device to a Device
func fromBusinessDevice(d device.Device) Device {
	return Device{
		ID:        d.ID,
		Name:      d.Name,
		Brand:     d.Brand,
		CreatedAt: d.CreatedAt,
	}
}
