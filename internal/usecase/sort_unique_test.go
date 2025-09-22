package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		field    int
		expected string
	}{
		{
			name:     "unique by whole line",
			input:    "apple\nbanana\napple\ncherry",
			field:    0,
			expected: "apple\nbanana\ncherry",
		},
		{
			name:     "unique by field 1",
			input:    "apple red\nbanana yellow\napple green\ncherry red",
			field:    1,
			expected: "apple red\nbanana yellow\ncherry red",
		},
		{
			name:     "unique by field 2",
			input:    "apple red\nbanana yellow\ngrape red\ncherry blue",
			field:    2,
			expected: "apple red\nbanana yellow\ncherry blue",
		},
		{
			name:     "no duplicates",
			input:    "apple\nbanana\ncherry",
			field:    0,
			expected: "apple\nbanana\ncherry",
		},
		{
			name:     "all duplicates",
			input:    "apple\napple\napple",
			field:    0,
			expected: "apple",
		},
		{
			name:     "empty input",
			input:    "",
			field:    0,
			expected: "",
		},
		{
			name:     "single line",
			input:    "apple",
			field:    0,
			expected: "apple",
		},
		{
			name:     "field beyond available",
			input:    "apple\nbanana yellow\ncherry red blue",
			field:    3,
			expected: "apple\ncherry red blue",
		},
		{
			name:     "mixed lines with and without field",
			input:    "apple red\nbanana\ngrape red\ncherry",
			field:    2,
			expected: "apple red\nbanana",
		},
		{
			name:     "empty lines",
			input:    "apple\n\nbanana\n\ncherry",
			field:    0,
			expected: "apple\n\nbanana\ncherry",
		},
		{
			name:     "whitespace lines",
			input:    "apple\n  \nbanana\n  \ncherry",
			field:    0,
			expected: "apple\n  \nbanana\ncherry",
		},
		{
			name:     "unique by field with empty key",
			input:    "apple\nbanana yellow\ncherry\ngrape green",
			field:    2,
			expected: "apple\nbanana yellow\ngrape green",
		},
		{
			name:     "preserve order",
			input:    "zebra\napple\nbanana\napple\nzebra",
			field:    0,
			expected: "zebra\napple\nbanana",
		},
		{
			name:     "case sensitive",
			input:    "Apple\napple\nAPPLE",
			field:    0,
			expected: "Apple\napple\nAPPLE",
		},
		{
			name:     "numbers as strings",
			input:    "1\n2\n1\n3\n2",
			field:    0,
			expected: "1\n2\n3",
		},
		{
			name:     "field with spaces",
			input:    "a b c\nx y z\na b d\nm n o",
			field:    2,
			expected: "a b c\nx y z\nm n o",
		},
		{
			name:     "multiple spaces between fields",
			input:    "a   b\nc   d\na   e\nf   b",
			field:    2,
			expected: "a   b\nc   d\na   e",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Unique(tt.input, tt.field)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestUniqueEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		field    int
		expected string
	}{
		{
			name:     "negative field defaults to whole line",
			input:    "apple\nbanana\napple",
			field:    -1,
			expected: "apple\nbanana",
		},
		{
			name:     "zero field defaults to whole line",
			input:    "apple\nbanana\napple",
			field:    0,
			expected: "apple\nbanana",
		},
		{
			name:     "field 1 same as whole line for single words",
			input:    "apple\nbanana\napple",
			field:    1,
			expected: "apple\nbanana",
		},
		{
			name:     "large field number",
			input:    "a b c\nx y z\na b c",
			field:    100,
			expected: "a b c",
		},
		{
			name:     "tabs as separators",
			input:    "a\tb\nc\td\na\te",
			field:    1,
			expected: "a\tb\nc\td",
		},
		{
			name:     "mixed separators",
			input:    "a b\tc\nx  y\tz\na\tb c",
			field:    2,
			expected: "a b\tc\nx  y\tz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Unique(tt.input, tt.field)
			require.Equal(t, tt.expected, result)
		})
	}
}
