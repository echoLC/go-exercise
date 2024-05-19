package exercise_test

import (
	"go-exercise/exercise"
	"testing"
)

func TestToRNA (t *testing.T) {
	var testCases = []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "Empty RNA sequence",
			input:       "",
			expected:    "",
		},
		{
			description: "RNA complement of cytosine is guanine",
			input:       "C",
			expected:    "G",
		},
		{
			description: "RNA complement of guanine is cytosine",
			input:       "G",
			expected:    "C",
		},
		{
			description: "RNA complement of thymine is adenine",
			input:       "T",
			expected:    "A",
		},
		{
			description: "RNA complement of adenine is uracil",
			input:       "A",
			expected:    "U",
		},
		{
			description: "RNA complement",
			input:       "ACGTGGTCTTAA",
			expected:    "UGCACCAGAAUU",
		},
	}

	for _, c := range testCases {
		got := exercise.ToRNA(c.input)
		if got != c.expected {
				t.Errorf("ReverseRunes(%q) == %q, want %q", c.input, got, c.expected)
		}
}
}