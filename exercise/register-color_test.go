package exercise_test

import (
	"go-exercise/exercise"
	"testing"
)

func TestRegisterColor (t *testing.T) {
	var testCases = []struct {
		description string
		input       []string
		expected    int
	}{
		{
			description: "Brown and black",
			input:       []string{"brown", "black"},
			expected:    10,
		},
		{
			description: "Blue and grey",
			input:       []string{"blue", "grey"},
			expected:    68,
		},
		{
			description: "Yellow and violet",
			input:       []string{"yellow", "violet"},
			expected:    47,
		},
		{
			description: "White and red",
			input:       []string{"white", "red"},
			expected:    92,
		},
		{
			description: "Orange and orange",
			input:       []string{"orange", "orange"},
			expected:    33,
		},
		{
			description: "Ignore additional colors",
			input:       []string{"green", "brown", "orange"},
			expected:    51,
		},
		{
			description: "Black and brown, one-digit",
			input:       []string{"black", "brown"},
			expected:    1,
		},
	}

	for _, c := range testCases {
		got := exercise.RegisterColor(c.input)
		if got != c.expected {
				t.Errorf("RegisterColor(%+q) == %d, want %d", c.input, got, c.expected)
		}
	}
}