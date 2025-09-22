package usecase

import (
	"testing"
	"unix_sort_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestSortByHumanNumeric(t *testing.T) {
	identity := func(s string) string { return s }

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic human numeric sort",
			input:    "1M\n2K\n500",
			expected: "500\n2K\n1M",
		},
		{
			name:     "negative numbers",
			input:    "-2M\n-1K\n-500",
			expected: "-2M\n-1K\n-500",
		},
		{
			name:     "mixed positive and negative",
			input:    "1M\n-2K\n500\n-1M",
			expected: "-1M\n-2K\n500\n1M",
		},
		{
			name:     "decimal numbers",
			input:    "1.5K\n2K\n1K",
			expected: "1K\n1.5K\n2K",
		},
		{
			name:     "same suffix different numbers",
			input:    "10K\n2K\n1K",
			expected: "1K\n2K\n10K",
		},
		{
			name:     "different suffixes same numbers",
			input:    "1G\n1M\n1K",
			expected: "1K\n1M\n1G",
		},
		{
			name:     "numbers without suffix",
			input:    "1000\n500\n2000",
			expected: "500\n1000\n2000",
		},
		{
			name:     "mixed with and without suffix",
			input:    "1K\n500\n2K\n1500",
			expected: "500\n1500\n1K\n2K",
		},
		{
			name:     "non-numeric strings",
			input:    "abc\n1K\nxyz\n500",
			expected: "500\n1K\nabc\nxyz",
		},
		{
			name:     "all suffixes order",
			input:    "1Q\n1R\n1Y\n1Z\n1E\n1P\n1T\n1G\n1M\n1K",
			expected: "1K\n1M\n1G\n1T\n1P\n1E\n1Z\n1Y\n1R\n1Q",
		},
		{
			name:     "case insensitive suffixes",
			input:    "1k\n1M\n1g",
			expected: "1k\n1M\n1g",
		},
		{
			name:     "zero values",
			input:    "0K\n0\n0M",
			expected: "0\n0K\n0M",
		},
		{
			name:     "negative with different suffixes",
			input:    "-1M\n-2K\n-1K",
			expected: "-1M\n-2K\n-1K",
		},
		{
			name:     "complex negative case",
			input:    "-2M\n-1M\n-2K\n-1K\n-2\n-1",
			expected: "-2M\n-1M\n-2K\n-1K\n-2\n-1",
		},
		{
			name:     "full range positive to negative",
			input:    "2M\n1K\n-1K\n-2M\n0\n500",
			expected: "-2M\n-1K\n0\n500\n1K\n2M",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "single line",
			input:    "1K",
			expected: "1K",
		},
		{
			name:     "whitespace handling",
			input:    "  1K  \n 2M \n500",
			expected: "500\n  1K  \n 2M ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByHumanNumeric(tt.input, identity, domain.SortOptions{})
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCompareHumanNumericStrings(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected bool
	}{
		{
			name:     "1K < 1M",
			a:        "1K",
			b:        "1M",
			expected: true,
		},
		{
			name:     "2K < 1M",
			a:        "2K",
			b:        "1M",
			expected: true,
		},
		{
			name:     "1K < 2K",
			a:        "1K",
			b:        "2K",
			expected: true,
		},
		{
			name:     "500 < 1K",
			a:        "500",
			b:        "1K",
			expected: true,
		},
		{
			name:     "-2K < -1K",
			a:        "-2K",
			b:        "-1K",
			expected: true,
		},
		{
			name:     "-1M < -1K",
			a:        "-1M",
			b:        "-1K",
			expected: true,
		},
		{
			name:     "1K > abc",
			a:        "1K",
			b:        "abc",
			expected: true,
		},
		{
			name:     "abc < xyz",
			a:        "abc",
			b:        "xyz",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := compareHumanNumericStrings(tt.a, tt.b)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSortByHumanNumericWithModify(t *testing.T) {
	trimBlanks := func(s string) string {
		return ignoreTrailingBlanks(s)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trailing spaces in human numeric",
			input:    "1K  \n2M \n500",
			expected: "500\n1K  \n2M ",
		},
		{
			name:     "trailing tabs in human numeric",
			input:    "1M\t\n2K\t\n500",
			expected: "500\n2K\t\n1M\t",
		},
		{
			name:     "mixed trailing blanks",
			input:    "1G \t\n1M  \n1K\t ",
			expected: "1K\t \n1M  \n1G \t",
		},
		{
			name:     "negative numbers with trailing blanks",
			input:    "-1K  \n-2M \n-500",
			expected: "-2M \n-1K  \n-500",
		},
		{
			name:     "decimal with trailing blanks",
			input:    "1.5K \n2.0K\t\n1.0K  ",
			expected: "1.0K  \n1.5K \n2.0K\t",
		},
		{
			name:     "non-numeric with trailing blanks",
			input:    "abc  \n1K \nxyz\t",
			expected: "1K \nabc  \nxyz\t",
		},
		{
			name:     "whitespace only lines",
			input:    "1K\n  \t  \n2M",
			expected: "1K\n2M\n  \t  ",
		},
		{
			name:     "preserve original formatting",
			input:    "  1K  \n 2M \n500",
			expected: "500\n  1K  \n 2M ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByHumanNumeric(tt.input, trimBlanks, domain.SortOptions{})
			require.Equal(t, tt.expected, result)
		})
	}
}
