package service

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	service := NewUserService()

	user, err := service.CreateUser("John Doe", "john@example.com", "password123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID == "" {
		t.Error("user ID should not be empty")
	}
	if user.Name != "John Doe" {
		t.Errorf("expected name %q, got %q", "John Doe", user.Name)
	}
	if user.Email != "john@example.com" {
		t.Errorf("expected email %q, got %q", "john@example.com", user.Email)
	}
}

func TestCreateUserValidation(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		email     string
		password  string
		wantError string
	}{
		{"empty name", "", "john@example.com", "password123", "name is required"},
		{"invalid email", "John Doe", "invalid-email", "password123", "invalid email format"},
		{"short password", "John Doe", "john@example.com", "short", "password must be at least 8 characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewUserService()
			_, err := service.CreateUser(tt.username, tt.email, tt.password)
			if err == nil {
				t.Fatalf("expected error %q, got nil", tt.wantError)
			}
		})
	}
}

func TestCreateUserDuplicateEmail(t *testing.T) {
	service := NewUserService()

	service.CreateUser("John Doe", "john@example.com", "password123")
	_, err := service.CreateUser("Jane Doe", "john@example.com", "password456")

	if err == nil {
		t.Error("expected error for duplicate email")
	}
}

func TestGetUserByID(t *testing.T) {
	service := NewUserService()
	created, _ := service.CreateUser("John Doe", "john@example.com", "password123")

	user := service.GetUserByID(created.ID)
	if user == nil {
		t.Error("expected to find user")
	}
	if user.ID != created.ID {
		t.Errorf("expected ID %q, got %q", created.ID, user.ID)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	service := NewUserService()
	user := service.GetUserByID("nonexistent")
	if user != nil {
		t.Error("expected nil for non-existent user")
	}
}

func TestGetAllUsers(t *testing.T) {
	service := NewUserService()

	if len(service.GetAllUsers()) != 0 {
		t.Error("expected empty user list initially")
	}

	service.CreateUser("John Doe", "john@example.com", "password123")
	service.CreateUser("Jane Doe", "jane@example.com", "password456")

	users := service.GetAllUsers()
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestUpdateUser(t *testing.T) {
	service := NewUserService()
	created, _ := service.CreateUser("John Doe", "john@example.com", "password123")

	updates := map[string]interface{}{"name": "Jane Doe"}
	user, err := service.UpdateUser(created.ID, updates)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Name != "Jane Doe" {
		t.Errorf("expected name %q, got %q", "Jane Doe", user.Name)
	}
}

func TestUpdateUserNotFound(t *testing.T) {
	service := NewUserService()
	_, err := service.UpdateUser("nonexistent", map[string]interface{}{})
	if err == nil {
		t.Error("expected error for non-existent user")
	}
}

func TestDeleteUser(t *testing.T) {
	service := NewUserService()
	created, _ := service.CreateUser("John Doe", "john@example.com", "password123")

	err := service.DeleteUser(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	user := service.GetUserByID(created.ID)
	if user != nil {
		t.Error("expected user to be deleted")
	}
}

func TestDeleteUserNotFound(t *testing.T) {
	service := NewUserService()
	err := service.DeleteUser("nonexistent")
	if err == nil {
		t.Error("expected error for non-existent user")
	}
}
