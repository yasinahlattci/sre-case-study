package handler

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetItem(ctx context.Context, objectID string) (interface{}, error) {
	args := m.Called(ctx, objectID)
	return args.Get(0), args.Error(1)
}

func (m *MockService) PutItem(ctx context.Context, objectID string, data interface{}) error {
	args := m.Called(ctx, objectID, data)
	return args.Error(0)
}

func (m *MockService) ListItems(ctx context.Context) ([]map[string]interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func TestHandler_GetItemHandler(t *testing.T) {
	tests := []struct {
		name       string
		objectID   string
		response   interface{}
		err        error
		wantStatus int
	}{
		{
			name:       "success",
			objectID:   "1",
			response:   map[string]interface{}{"age": 28, "city": "Denizli", "name": "Yasin"},
			err:        nil,
			wantStatus: 200,
		},
		{
			name:       "not found",
			objectID:   "2",
			response:   nil,
			err:        nil,
			wantStatus: 404,
		},
		{
			name:       "internal error",
			objectID:   "3",
			response:   nil,
			err:        errors.New("db error"),
			wantStatus: 500,
		},
		{
			name:       "timeout",
			objectID:   "4",
			response:   nil,
			err:        context.DeadlineExceeded,
			wantStatus: 408,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			mockService := new(MockService)
			h := &Handler{dbService: mockService}

			mockService.On("GetItem", mock.Anything, tt.objectID).Return(tt.response, tt.err)

			app.Get("/picus/get/:objectID", h.GetItemHandler)
			req := httptest.NewRequest("GET", "/picus/get/"+tt.objectID, nil)
			resp, _ := app.Test(req)

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}

func TestHandler_ListItemsHandler(t *testing.T) {
	tests := []struct {
		name       string
		response   []map[string]interface{}
		err        error
		wantStatus int
	}{
		{
			name:       "success",
			response:   []map[string]interface{}{{"data": map[string]interface{}{"age": 28, "city": "Denizli", "name": "Yasin"}, "objectID": "notimportant"}},
			err:        nil,
			wantStatus: 200,
		},
		{
			name:       "success-2",
			response:   []map[string]interface{}{{"data": map[string]interface{}{"age": 28, "city": "Denizli", "name": "Yasin"}, "objectID": "notimportant"}, {"data": map[string]interface{}{"age": 30, "city": "Istanbul", "name": "Ahmet"}, "objectID": "notimportant1"}},
			err:        nil,
			wantStatus: 200,
		},
		{
			name:       "internal error",
			response:   nil,
			err:        errors.New("db error"),
			wantStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			mockService := new(MockService)
			h := &Handler{dbService: mockService}

			mockService.On("ListItems", mock.Anything).Return(tt.response, tt.err)

			app.Get("/picus/list", h.ListItemsHandler)
			req := httptest.NewRequest("GET", "/picus/list", nil)
			resp, _ := app.Test(req)

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}


func TestHandler_PutItemHandler(t *testing.T) {
	tests := []struct {
		name       string
		requestBody string
		err        error
		wantStatus int
	}{
		{
			name:       "success",
			requestBody: `{"data": {"age": 28, "city": "Denizli", "name": "Yasin"}}`,
			err:        nil,
			wantStatus: 201,
		},
		{
			name:       "invalid request body",
			requestBody: `invalid json`,
			err:        nil,
			wantStatus: 400,
		},
		{
			name:       "missing data field in request body",
			requestBody: `{}`,
			err:        nil,
			wantStatus: 400,
		},
		{
			name:       "internal error",
			requestBody: `{"data": {"age": 28, "city": "Denizli", "name": "Yasin"}}`,
			err:        errors.New("db error"),
			wantStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			mockService := new(MockService)
			h := &Handler{dbService: mockService}


			mockService.On("PutItem", mock.Anything, mock.Anything, mock.Anything).Return(tt.err)

			app.Post("/picus/put", h.PutItemHandler)
			req := httptest.NewRequest("POST", "/picus/put", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}