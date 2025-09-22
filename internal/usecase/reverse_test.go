package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "multiple lines",
			input:    "a\nb\nc",
			expected: "c\nb\na",
		},
		{
			name:     "single line",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "two lines",
			input:    "first\nsecond",
			expected: "second\nfirst",
		},
		{
			name:     "even number of lines",
			input:    "1\n2\n3\n4",
			expected: "4\n3\n2\n1",
		},
		{
			name:     "odd number of lines",
			input:    "1\n2\n3\n4\n5",
			expected: "5\n4\n3\n2\n1",
		},
		{
			name:     "empty lines",
			input:    "a\n\nb",
			expected: "b\n\na",
		},
		{
			name:     "whitespace lines",
			input:    "first\n  \nlast",
			expected: "last\n  \nfirst",
		},
		{
			name:     "identical lines",
			input:    "same\nsame\nsame",
			expected: "same\nsame\nsame",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Reverse(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
