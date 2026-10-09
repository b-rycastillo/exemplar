package service

import (
	"errors"
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserService handles user business logic
type UserService struct {
	users []*User
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users: []*User{},
	}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if !IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	if !IsValidPassword(password) {
		return nil, errors.New("password must be at least 8 characters")
	}

	// Check if email already exists
	for _, u := range us.users {
		if u.Email == email {
			return nil, errors.New("email already exists")
		}
	}

	user := &User{
		ID:        GenerateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	us.users = append(us.users, user)
	return user, nil
}

// GetUserByID retrieves a user by ID
func (us *UserService) GetUserByID(id string) *User {
	for _, u := range us.users {
		if u.ID == id {
			return u
		}
	}
	return nil
}

// GetAllUsers returns all users
func (us *UserService) GetAllUsers() []*User {
	return us.users
}

// UpdateUser updates a user
func (us *UserService) UpdateUser(id string, updates map[string]interface{}) (*User, error) {
	user := us.GetUserByID(id)
	if user == nil {
		return nil, errors.New("user not found")
	}

	if newEmail, ok := updates["email"].(string); ok {
		if !IsValidEmail(newEmail) {
			return nil, errors.New("invalid email format")
		}
		// Check if new email already exists
		for _, u := range us.users {
			if u.Email == newEmail && u.ID != id {
				return nil, errors.New("email already exists")
			}
		}
		user.Email = newEmail
	}

	if newName, ok := updates["name"].(string); ok {
		user.Name = newName
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(id string) error {
	for i, u := range us.users {
		if u.ID == id {
			us.users = append(us.users[:i], us.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
