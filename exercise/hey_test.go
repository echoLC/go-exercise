package exercise_test

import (
	"go-exercise/exercise"
	"testing"
)

func TestHey (t *testing.T) {
	var testCases = []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "stating something",
			input:       "Tom-ay-to, tom-aaaah-to.",
			expected:    "Whatever.",
		},
		{
			description: "shouting",
			input:       "WATCH OUT!",
			expected:    "Whoa, chill out!",
		},
		{
			description: "shouting gibberish",
			input:       "FCECDFCAAB",
			expected:    "Whoa, chill out!",
		},
		{
			description: "asking a question",
			input:       "Does this cryogenic chamber make me look fat?",
			expected:    "Sure.",
		},
		{
			description: "asking a numeric question",
			input:       "You are, what, like 15?",
			expected:    "Sure.",
		},
		{
			description: "asking gibberish",
			input:       "fffbbcbeab?",
			expected:    "Sure.",
		},
		{
			description: "talking forcefully",
			input:       "Hi there!",
			expected:    "Whatever.",
		},
		{
			description: "using acronyms in regular speech",
			input:       "It's OK if you don't want to go work for NASA.",
			expected:    "Whatever.",
		},
		{
			description: "forceful question",
			input:       "WHAT'S GOING ON?",
			expected:    "Calm down, I know what I'm doing!",
		},
		{
			description: "shouting numbers",
			input:       "1, 2, 3 GO!",
			expected:    "Whoa, chill out!",
		},
		{
			description: "no letters",
			input:       "1, 2, 3",
			expected:    "Whatever.",
		},
		{
			description: "question with no letters",
			input:       "4?",
			expected:    "Sure.",
		},
		{
			description: "shouting with special characters",
			input:       "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
			expected:    "Whoa, chill out!",
		},
		{
			description: "shouting with no exclamation mark",
			input:       "I HATE THE DENTIST",
			expected:    "Whoa, chill out!",
		},
		{
			description: "statement containing question mark",
			input:       "Ending with ? means a question.",
			expected:    "Whatever.",
		},
		{
			description: "non-letters with question",
			input:       ":) ?",
			expected:    "Sure.",
		},
		{
			description: "prattling on",
			input:       "Wait! Hang on. Are you going to be OK?",
			expected:    "Sure.",
		},
		{
			description: "silence",
			input:       "",
			expected:    "Fine. Be that way!",
		},
		{
			description: "prolonged silence",
			input:       "          ",
			expected:    "Fine. Be that way!",
		},
		{
			description: "alternate silence",
			input:       "\t\t\t\t\t\t\t\t\t\t",
			expected:    "Fine. Be that way!",
		},
		{
			description: "multiple line question",
			input:       "\nDoes this cryogenic chamber make me look fat?\nNo.",
			expected:    "Whatever.",
		},
		{
			description: "starting with whitespace",
			input:       "         hmmmmmmm...",
			expected:    "Whatever.",
		},
		{
			description: "ending with whitespace",
			input:       "Okay if like my  spacebar  quite a bit?   ",
			expected:    "Sure.",
		},
		{
			description: "other whitespace",
			input:       "\n\r \t",
			expected:    "Fine. Be that way!",
		},
		{
			description: "non-question ending with whitespace",
			input:       "This is a statement ending with whitespace      ",
			expected:    "Whatever.",
		},
	}

	for _, c := range testCases {
		got := exercise.Hey(c.input)
		if got != c.expected {
				t.Errorf("ReverseRunes(%q) == %q, want %q", c.input, got, c.expected)
		}
	}
}