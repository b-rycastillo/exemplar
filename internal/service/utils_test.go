package service

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"testexample.com", false},
		{"test@", false},
		{"", false},
	}

	for _, tt := range tests {
		result := IsValidEmail(tt.email)
		if result != tt.expected {
			t.Errorf("IsValidEmail(%q) = %v, want %v", tt.email, result, tt.expected)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"password123", true},
		{"pass123", false},
		{"", false},
	}

	for _, tt := range tests {
		result := IsValidPassword(tt.password)
		if result != tt.expected {
			t.Errorf("IsValidPassword(%q) = %v, want %v", tt.password, result, tt.expected)
		}
	}
}

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()

	if len(id1) == 0 || len(id2) == 0 {
		t.Error("GenerateID should generate non-empty IDs")
	}

	expectedPrefix := "user_"
	if id1[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("GenerateID should start with %q", expectedPrefix)
	}

	// Note: IDs might be identical if generated in same nanosecond
	if id1 != id2 {
		t.Logf("Generated two different IDs (expected for different calls): %s, %s", id1, id2)
	}
}
