package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/server"
	"github.com/ss49919201/go-functional-layered-architecture/in-memory/internal/service"
)

func setupTestServer() http.Handler {
	return server.NewHandler()
}

func TestRetriveReservation_Success(t *testing.T) {
	server := setupTestServer()

	req := httptest.NewRequest("GET", "/reservation/1", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type 'application/json', got '%s'", contentType)
	}

	var result service.RetriveReservationResult
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedID := 1
	if result.ID != expectedID {
		t.Errorf("Expected ID %d, got %d", expectedID, result.ID)
	}

	expectedRoomName := "normal"
	if result.RoomName != expectedRoomName {
		t.Errorf("Expected RoomName '%s', got '%s'", expectedRoomName, result.RoomName)
	}

	expectedReservationDateTime := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	if !result.ReservationDateTime.Equal(expectedReservationDateTime) {
		t.Errorf("Expected ReservationDateTime %v, got %v", expectedReservationDateTime, result.ReservationDateTime)
	}

	expectedConfirmedDateTime := time.Date(2024, 1, 1, 5, 0, 0, 0, time.UTC)
	if !result.ConfirmedDateTime.Equal(expectedConfirmedDateTime) {
		t.Errorf("Expected ConfirmedDateTime %v, got %v", expectedConfirmedDateTime, result.ConfirmedDateTime)
	}
}

func TestRetriveReservation_NotFound(t *testing.T) {
	server := setupTestServer()

	req := httptest.NewRequest("GET", "/reservation/999", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	expectedBody := "reservation not found\n"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestRetriveReservation_InvalidID(t *testing.T) {
	server := setupTestServer()

	req := httptest.NewRequest("GET", "/reservation/invalid", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	expectedBody := "invalid id parameter\n"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
	}
}

func TestRetriveReservation_MultipleReservations(t *testing.T) {
	server := setupTestServer()

	testCases := []struct {
		id               string
		expectedRoomName string
		expectedStatus   int
	}{
		{"1", "normal", http.StatusOK},
		{"2", "normal", http.StatusOK},
		{"3", "special", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run("ID_"+tc.id, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/reservation/"+tc.id, nil)
			w := httptest.NewRecorder()

			server.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.expectedStatus, w.Code)
			}

			if tc.expectedStatus == http.StatusOK {
				var result service.RetriveReservationResult
				err := json.Unmarshal(w.Body.Bytes(), &result)
				if err != nil {
					t.Fatalf("Failed to unmarshal response: %v", err)
				}

				if result.RoomName != tc.expectedRoomName {
					t.Errorf("Expected RoomName '%s', got '%s'", tc.expectedRoomName, result.RoomName)
				}
			}
		})
	}
}
