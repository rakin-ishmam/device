package mocks

import (
	"context"
	"device/business/device"

	"github.com/stretchr/testify/mock"
)

// Store is a mock type for the store
type Store struct {
	mock.Mock
}

var _ device.Store = (*Store)(nil)

func (s *Store) ByID(ctx context.Context, id string) (device.Device, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(device.Device), args.Error(1)
}

func (s *Store) Create(ctx context.Context, d device.Device) error {
	args := s.Called(ctx, d)
	return args.Error(0)
}

func (s *Store) Update(ctx context.Context, id string, data device.UpdateDevice) error {
	args := s.Called(ctx, id, data)
	return args.Error(0)
}

func (s *Store) GetAll(ctx context.Context, offset, limit int) ([]device.Device, error) {
	args := s.Called(ctx, offset, limit)
	return args.Get(0).([]device.Device), args.Error(1)
}

func (s *Store) Delete(ctx context.Context, id string) error {
	args := s.Called(ctx, id)
	return args.Error(0)
}

func (s *Store) SearchByBrand(ctx context.Context, brand string, offset, limit int) ([]device.Device, error) {
	args := s.Called(ctx, brand, offset, limit)
	return args.Get(0).([]device.Device), args.Error(1)
}
