package tests

import (
	"backend/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		name          string
		status        int
		data          interface{}
		expectedBody  string
		expectedError string
	}{
		{
			name:   "Valid data",
			status: http.StatusOK,
			data: utils.JSONResponse{
				Error:   false,
				Message: "Success",
				Data:    "Test",
			},
			expectedBody:  `{"error":false,"message":"Success","data":"Test"}`,
			expectedError: "",
		},
		{
			name:          "Nil data",
			status:        http.StatusOK,
			data:          nil,
			expectedBody:  "null",
			expectedError: "",
		},
		{
			name:          "Unserializable data",
			status:        http.StatusOK,
			data:          make(chan int),
			expectedBody:  "",
			expectedError: "failed to marshal data: json: unsupported type: chan int",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			err := utils.WriteJSON(w, test.status, test.data)

			if test.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if !strings.Contains(err.Error(), test.expectedError) {
					t.Errorf("Expected error '%s', but got '%s'", test.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if w.Body.String() != test.expectedBody {
				t.Errorf("Expected body '%s', but got '%s'", test.expectedBody, w.Body.String())
			}
		})
	}
}

func TestReadJSON(t *testing.T) {
	tests := []struct {
		name          string
		body          string
		expectedData  interface{}
		expectedError string
	}{
		{
			name: "Valid JSON",
			body: `{"name":"Test"}`,
			expectedData: map[string]interface{}{
				"name": "Test",
			},
			expectedError: "",
		},
		{
			name:          "Invalid JSON",
			body:          `{"name":"Test",}`,
			expectedData:  nil,
			expectedError: "invalid character '}' looking for beginning of object key string",
		},
		{
			name:          "Extra JSON object",
			body:          `{"name":"Test"}{}`,
			expectedData:  nil,
			expectedError: "body must only contain a single JSON value",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", bytes.NewBufferString(test.body))

			var data map[string]interface{}

			err := utils.ReadJSON(w, r, &data)

			if test.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if !strings.Contains(err.Error(), test.expectedError) {
					t.Errorf("Expected error '%s', but got '%s'", test.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if !jsonEqual(data, test.expectedData) {
				t.Errorf("Expected data %v, but got %v", test.expectedData, data)
			}
		})
	}
}

func TestErrorJSON(t *testing.T) {
	tests := []struct {
		name          string
		err           error
		status        []int
		expectedBody  string
		expectedError string
	}{
		{
			name:          "Error with default status",
			err:           errors.New("test error"),
			status:        []int{},
			expectedBody:  `{"error":true,"message":"test error"}`,
			expectedError: "",
		},
		{
			name:          "Error with custom status",
			err:           errors.New("test error"),
			status:        []int{http.StatusNotFound},
			expectedBody:  `{"error":true,"message":"test error"}`,
			expectedError: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			err := utils.ErrorJSON(w, test.err, test.status...)

			if test.expectedError != "" {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if !strings.Contains(err.Error(), test.expectedError) {
					t.Errorf("Expected error '%s', but got '%s'", test.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			// Remove status from body for comparison
			body := strings.TrimPrefix(w.Body.String(), fmt.Sprintf("%d ", w.Code))

			if body != test.expectedBody {
				t.Errorf("Expected body '%s', but got '%s'", test.expectedBody, body)
			}
		})
	}
}

func jsonEqual(a, b interface{}) bool {
	bytesA, _ := json.Marshal(a)
	bytesB, _ := json.Marshal(b)

	return bytes.Equal(bytesA, bytesB)
}
