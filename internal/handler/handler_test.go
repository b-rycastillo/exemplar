package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/example/exemplar/internal/service"
)

func TestHealthCheck(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp HealthResponse
	json.NewDecoder(w.Body).Decode(&resp)
	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got %q", resp.Status)
	}
}

func TestCreateUserAPI(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	body := CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	var user service.User
	json.NewDecoder(w.Body).Decode(&user)
	if user.Name != "John Doe" || user.Email != "john@example.com" {
		t.Error("user data mismatch")
	}
}

func TestCreateUserInvalidEmail(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	body := CreateUserRequest{
		Name:     "John Doe",
		Email:    "invalid-email",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateUserShortPassword(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	body := CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "short",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateUserDuplicateEmailAPI(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	body := CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Try creating again
	req = httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGetAllUsersAPI(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var users []*service.User
	json.NewDecoder(w.Body).Decode(&users)
	if len(users) != 0 {
		t.Errorf("expected 0 users initially, got %d", len(users))
	}
}

func TestGetUserByIDAPI(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	// Create a user
	body := CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var user service.User
	json.NewDecoder(w.Body).Decode(&user)
	userID := user.ID

	// Get the user
	req = httptest.NewRequest("GET", "/users/"+userID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	json.NewDecoder(w.Body).Decode(&user)
	if user.Name != "John Doe" {
		t.Errorf("expected name 'John Doe', got %q", user.Name)
	}
}

func TestGetUserNotFound(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	req := httptest.NewRequest("GET", "/users/nonexistent", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestUpdateUserAPI(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	// Create a user
	createBody := CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(createBody)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var user service.User
	json.NewDecoder(w.Body).Decode(&user)
	userID := user.ID

	// Update the user
	updateBody := UpdateUserRequest{
		Name: "Jane Doe",
	}
	bodyBytes, _ = json.Marshal(updateBody)

	req = httptest.NewRequest("PATCH", "/users/"+userID, bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	json.NewDecoder(w.Body).Decode(&user)
	if user.Name != "Jane Doe" {
		t.Errorf("expected name 'Jane Doe', got %q", user.Name)
	}
}

func TestDeleteUserAPI(t *testing.T) {
	svc := service.NewUserService()
	handler := NewHandler(svc)
	router := SetupRoutes(handler)

	// Create a user
	body := CreateUserRequest{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var user service.User
	json.NewDecoder(w.Body).Decode(&user)
	userID := user.ID

	// Delete the user
	req = httptest.NewRequest("DELETE", "/users/"+userID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Verify deletion
	req = httptest.NewRequest("GET", "/users/"+userID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d after deletion, got %d", http.StatusNotFound, w.Code)
	}
}
