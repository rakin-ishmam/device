package device

import (
	"bytes"
	"context"
	"device/business/device"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
)

type BusinessMock struct {
	mock.Mock
}

func (bm *BusinessMock) Create(ctx context.Context, cd device.CreateDevice) (device.Device, error) {
	args := bm.Called(ctx, cd)
	return args.Get(0).(device.Device), args.Error(1)
}

func (bm *BusinessMock) Update(ctx context.Context, id string, data device.UpdateDevice) error {
	args := bm.Called(ctx, id, data)
	return args.Error(0)
}

func (bm *BusinessMock) GetByID(ctx context.Context, id string) (device.Device, error) {
	args := bm.Called(ctx, id)
	return args.Get(0).(device.Device), args.Error(1)
}

func (bm *BusinessMock) GetAll(ctx context.Context, offset, limit int) ([]device.Device, error) {
	args := bm.Called(ctx, offset, limit)
	return args.Get(0).([]device.Device), args.Error(1)
}

func (bm *BusinessMock) Delete(ctx context.Context, id string) error {
	args := bm.Called(ctx, id)
	return args.Error(0)
}

func (bm *BusinessMock) SearchByBrand(ctx context.Context, brand string, offset, limit int) ([]device.Device, error) {
	args := bm.Called(ctx, brand, offset, limit)
	return args.Get(0).([]device.Device), args.Error(1)
}

func TestCreate(t *testing.T) {
	testTable := map[string]struct {
		request        device.CreateDevice
		expectedStatus int
	}{
		"success": {
			request: device.CreateDevice{
				Name:  "test",
				Brand: "test",
			},
			expectedStatus: http.StatusCreated,
		},
		"invalid request": {
			request: device.CreateDevice{
				Name:  "",
				Brand: "test",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for tn, tc := range testTable {
		t.Run(tn, func(t *testing.T) {
			b := BusinessMock{}
			b.On("Create", mock.Anything, mock.AnythingOfType("device.CreateDevice")).Return(device.Device{}, nil)
			h := NewHandler(&b)

			r := chi.NewRouter()
			r.Post("/device", h.Create)

			data, err := json.Marshal(tc.request)
			if err != nil {
				t.Fatalf("unable to marshal request: %v", err)
			}

			buf := bytes.NewBuffer(data)
			req := httptest.NewRequest("POST", "/device", buf)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != tc.expectedStatus {
				t.Fatalf("expected %d, got %d", rec.Code, tc.expectedStatus)
			}
		})
	}
}
