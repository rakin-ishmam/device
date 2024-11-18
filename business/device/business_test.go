package device_test

import (
	"context"
	"device/business/device"
	"device/business/device/store/mocks"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	testTable := map[string]struct {
		createDevice   device.CreateDevice
		expectedDevice device.Device
		dbResponse     error
		expectedErr    error
	}{
		"ok": {
			createDevice: device.CreateDevice{
				Name:  "name",
				Brand: "brand",
			},
			expectedDevice: device.Device{
				Name:  "name",
				Brand: "brand",
			},
			dbResponse:  nil,
			expectedErr: nil,
		},
		"error": {
			createDevice: device.CreateDevice{
				Name:  "name",
				Brand: "brand",
			},
			expectedDevice: device.Device{
				Name:  "name",
				Brand: "brand",
			},
			dbResponse:  fmt.Errorf("db err"),
			expectedErr: fmt.Errorf("store.Create: db err"),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			m := mocks.Store{}
			m.On("Create", mock.Anything, mock.AnythingOfType("device.Device")).Return(tc.dbResponse)

			b := device.NewBusiness(&m)
			result, err := b.Create(context.Background(), tc.createDevice)
			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil {
				return
			}
			if result.Name != tc.expectedDevice.Name || result.Brand != tc.expectedDevice.Brand {
				t.Fatalf("expected %v, got %v", tc.expectedDevice, result)
			}

			if _, err := uuid.Parse(result.ID); err != nil {
				t.Fatalf("expected a valid UUID, got %v", result.ID)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	testTable := map[string]struct {
		id          string
		updateData  device.UpdateDevice
		dbResponse  error
		expectedErr error
	}{
		"ok": {
			id: "1",
			updateData: device.UpdateDevice{
				Name: strPtr("new name"),
			},
			dbResponse:  nil,
			expectedErr: nil,
		},
		"error": {
			id: "1",
			updateData: device.UpdateDevice{
				Name: strPtr("new name"),
			},
			dbResponse:  device.ErrNotFound,
			expectedErr: fmt.Errorf("store.Update: %w", device.ErrNotFound),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			m := mocks.Store{}
			m.On("Update", mock.Anything, tc.id, tc.updateData).Return(tc.dbResponse)

			b := device.NewBusiness(&m)
			err := b.Update(context.Background(), tc.id, tc.updateData)
			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	testTable := map[string]struct {
		id          string
		dbResponse  device.Device
		dbError     error
		expectedErr error
	}{
		"ok": {
			id: "1",
			dbResponse: device.Device{
				ID:        "1",
				Name:      "name",
				Brand:     "brand",
				CreatedAt: time.Now(),
			},
			dbError:     nil,
			expectedErr: nil,
		},
		"not found": {
			id:          "2",
			dbResponse:  device.Device{},
			dbError:     device.ErrNotFound,
			expectedErr: fmt.Errorf("store.ByID: %w", device.ErrNotFound),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			m := mocks.Store{}
			m.On("ByID", mock.Anything, tc.id).Return(tc.dbResponse, tc.dbError)

			b := device.NewBusiness(&m)
			result, err := b.GetByID(context.Background(), tc.id)
			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(result, tc.dbResponse) {
				t.Fatalf("expected %v, got %v", tc.dbResponse, result)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	testTable := map[string]struct {
		offset      int
		limit       int
		dbResponse  []device.Device
		dbError     error
		expectedErr error
	}{
		"ok": {
			offset: 0,
			limit:  10,
			dbResponse: []device.Device{
				{
					ID:        "1",
					Name:      "name1",
					Brand:     "brand1",
					CreatedAt: time.Now(),
				},
				{
					ID:        "2",
					Name:      "name2",
					Brand:     "brand2",
					CreatedAt: time.Now(),
				},
			},
			dbError:     nil,
			expectedErr: nil,
		},
		"error": {
			offset:      0,
			limit:       10,
			dbResponse:  nil,
			dbError:     fmt.Errorf("db err"),
			expectedErr: fmt.Errorf("store.GetAll: db err"),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			m := mocks.Store{}
			m.On("GetAll", mock.Anything, tc.offset, tc.limit).Return(tc.dbResponse, tc.dbError)

			b := device.NewBusiness(&m)
			result, err := b.GetAll(context.Background(), tc.offset, tc.limit)
			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(result, tc.dbResponse) {
				t.Fatalf("expected %v, got %v", tc.dbResponse, result)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testTable := map[string]struct {
		id          string
		dbResponse  error
		expectedErr error
	}{
		"ok": {
			id:          "1",
			dbResponse:  nil,
			expectedErr: nil,
		},
		"not found": {
			id:          "2",
			dbResponse:  device.ErrNotFound,
			expectedErr: fmt.Errorf("store.Delete: %w", device.ErrNotFound),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			m := mocks.Store{}
			m.On("Delete", mock.Anything, tc.id).Return(tc.dbResponse)

			b := device.NewBusiness(&m)
			err := b.Delete(context.Background(), tc.id)
			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
func TestSearchByBrand(t *testing.T) {
	testTable := map[string]struct {
		brand       string
		offset      int
		limit       int
		dbResponse  []device.Device
		dbError     error
		expectedErr error
	}{
		"ok": {
			brand:  "brand1",
			offset: 0,
			limit:  10,
			dbResponse: []device.Device{
				{
					ID:        "1",
					Name:      "name1",
					Brand:     "brand1",
					CreatedAt: time.Now(),
				},
				{
					ID:        "2",
					Name:      "name2",
					Brand:     "brand1",
					CreatedAt: time.Now(),
				},
			},
			dbError:     nil,
			expectedErr: nil,
		},
		"error": {
			brand:       "brand2",
			offset:      0,
			limit:       10,
			dbResponse:  nil,
			dbError:     fmt.Errorf("db err"),
			expectedErr: fmt.Errorf("store.SearchByBrand: db err"),
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			m := mocks.Store{}
			m.On("SearchByBrand", mock.Anything, tc.brand, tc.offset, tc.limit).Return(tc.dbResponse, tc.dbError)

			b := device.NewBusiness(&m)
			result, err := b.SearchByBrand(context.Background(), tc.brand, tc.offset, tc.limit)
			if (err == nil && tc.expectedErr != nil) || (err != nil && tc.expectedErr == nil) {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected %v, got %v", tc.expectedErr, err)
			}

			if !reflect.DeepEqual(result, tc.dbResponse) {
				t.Fatalf("expected %v, got %v", tc.dbResponse, result)
			}
		})
	}
}
