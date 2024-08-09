package services

import (
	"testing"
)

func TestCalculateStrongPasswordSteps(t *testing.T) {
	tests := []struct {
		password string
		expected int
	}{
		// Test cases with different passwords
		{"abc", 3}, // Less than 6 characters
		{"aBc2ef", 0}, // Valid password (length 6)
		{"aBc1", 2}, // Less than 6 characters, missing types
		{"aBcdef1234567890", 0}, // Valid password (length 16)
		{"aA1", 3}, // Less than 6 characters, missing types
		{"aaAA11", 0}, // Valid password (length 6) with some repeats
		{"aaAA1111", 2}, // Password with repeats
		{"aA1aaaAA", 1}, // Length is okay, but has repeats
		{"a234enleiFdkjqierj2n3o4j", 4}, // Password with length 23
	}

	for _, test := range tests {
		result := CalculateStrongPasswordSteps(test.password)
		if result != test.expected {
			t.Errorf("CalculateStrongPasswordSteps(%q) = %d; want %d", test.password, result, test.expected)
		}
	}
}