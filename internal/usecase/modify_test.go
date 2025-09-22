package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIgnoreTrailingBlanks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trailing spaces",
			input:    "hello  ",
			expected: "hello",
		},
		{
			name:     "trailing tabs",
			input:    "hello\t\t",
			expected: "hello",
		},
		{
			name:     "mixed trailing blanks",
			input:    "hello \t ",
			expected: "hello",
		},
		{
			name:     "no trailing blanks",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "leading spaces preserved",
			input:    "  hello  ",
			expected: "  hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only blanks",
			input:    " \t ",
			expected: "",
		},
		{
			name:     "single space",
			input:    "a ",
			expected: "a",
		},
		{
			name:     "single tab",
			input:    "a\t",
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ignoreTrailingBlanks(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
