package exercise_test

import (
	"go-exercise/exercise"
	"testing"
)

func TestGrains (t *testing.T) {
	var testCases = []struct {
		description string
		input       int
		expectedVal uint64
		expectError bool
	}{
		{
			description: "grains on square 1",
			input:       1,
			expectedVal: 1,
			expectError: false,
		},
		{
			description: "grains on square 2",
			input:       2,
			expectedVal: 2,
			expectError: false,
		},
		{
			description: "grains on square 3",
			input:       3,
			expectedVal: 4,
			expectError: false,
		},
		{
			description: "grains on square 4",
			input:       4,
			expectedVal: 8,
			expectError: false,
		},
		{
			description: "grains on square 16",
			input:       16,
			expectedVal: 32768,
			expectError: false,
		},
		{
			description: "grains on square 32",
			input:       32,
			expectedVal: 2147483648,
			expectError: false,
		},
		{
			description: "grains on square 64",
			input:       64,
			expectedVal: 9223372036854775808,
			expectError: false,
		},
		{
			description: "square 0 raises an exception",
			input:       0,
			expectedVal: 0,
			expectError: true,
		},
		{
			description: "negative square raises an exception",
			input:       -1,
			expectedVal: 0,
			expectError: true,
		},
		{
			description: "square greater than 64 raises an exception",
			input:       65,
			expectedVal: 0,
			expectError: true,
		},
	}

	for _, c := range testCases {
		got, _ := exercise.Square(c.input)
		if got != c.expectedVal {
				t.Errorf("ReverseRunes(%q) == %q, want %q", c.input, got, c.expectedVal)
		}
	}
}