package usecase

import (
	"testing"
	"unix_sort_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestSortByNumeric(t *testing.T) {
	identity := func(s string) string { return s }

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic numeric sort",
			input:    "10\n2\n1",
			expected: "1\n2\n10",
		},
		{
			name:     "negative numbers",
			input:    "-5\n-10\n-1",
			expected: "-10\n-5\n-1",
		},
		{
			name:     "mixed positive and negative",
			input:    "5\n-3\n0\n-10\n2",
			expected: "-10\n-3\n0\n2\n5",
		},
		{
			name:     "decimal numbers",
			input:    "1.5\n2.1\n1.2",
			expected: "1.2\n1.5\n2.1",
		},
		{
			name:     "mixed integers and decimals",
			input:    "10\n1.5\n2\n1.2",
			expected: "1.2\n1.5\n2\n10",
		},
		{
			name:     "numbers with text",
			input:    "file10.txt\nfile2.txt\nfile1.txt",
			expected: "file1.txt\nfile10.txt\nfile2.txt",
		},
		{
			name:     "numbers at different positions",
			input:    "version-10\nversion-2\nversion-1",
			expected: "version-1\nversion-10\nversion-2",
		},
		{
			name:     "non-numeric strings",
			input:    "abc\n5\nxyz\n1",
			expected: "1\n5\nabc\nxyz",
		},
		{
			name:     "zero values",
			input:    "0\n-0\n0.0",
			expected: "-0\n0\n0.0",
		},
		{
			name:     "large numbers",
			input:    "1000000\n999999\n1000001",
			expected: "999999\n1000000\n1000001",
		},
		{
			name:     "scientific notation not supported",
			input:    "1e3\n2000\n1000",
			expected: "1000\n2000\n1e3",
		},
		{
			name:     "numbers with plus sign",
			input:    "+5\n-3\n+1",
			expected: "-3\n+1\n+5",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "single number",
			input:    "42",
			expected: "42",
		},
		{
			name:     "only non-numeric",
			input:    "apple\nbanana\ncherry",
			expected: "apple\nbanana\ncherry",
		},
		{
			name:     "numbers in middle of text",
			input:    "test123end\ntest45end\ntest7end",
			expected: "test123end\ntest45end\ntest7end",
		},
		{
			name:     "multiple numbers in line",
			input:    "1 and 100\n2 and 50\n3 and 200",
			expected: "1 and 100\n2 and 50\n3 and 200",
		},
		{
			name:     "fractional numbers",
			input:    "0.1\n0.01\n0.001",
			expected: "0.001\n0.01\n0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByNumeric(tt.input, identity, domain.SortOptions{})
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCompareNumericStrings(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected bool
	}{
		{
			name:     "1 < 2",
			a:        "1",
			b:        "2",
			expected: true,
		},
		{
			name:     "10 > 2",
			a:        "10",
			b:        "2",
			expected: false,
		},
		{
			name:     "-5 < -1",
			a:        "-5",
			b:        "-1",
			expected: true,
		},
		{
			name:     "1.5 < 2",
			a:        "1.5",
			b:        "2",
			expected: true,
		},
		{
			name:     "number < non-number",
			a:        "5",
			b:        "abc",
			expected: true,
		},
		{
			name:     "non-number > number",
			a:        "abc",
			b:        "5",
			expected: false,
		},
		{
			name:     "non-number < non-number",
			a:        "abc",
			b:        "xyz",
			expected: true,
		},
		{
			name:     "zero comparison",
			a:        "0",
			b:        "-0",
			expected: false,
		},
		{
			name:     "decimal precision",
			a:        "1.1",
			b:        "1.10",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := compareNumericStrings(tt.a, tt.b)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSortByNumericWithModify(t *testing.T) {
	trimBlanks := func(s string) string {
		return ignoreTrailingBlanks(s)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "numbers with trailing spaces",
			input:    "10  \n2 \n1",
			expected: "1\n2 \n10  ",
		},
		{
			name:     "numbers with trailing tabs",
			input:    "10\t\n2\t\t\n1",
			expected: "1\n2\t\t\n10\t",
		},
		{
			name:     "mixed trailing blanks",
			input:    "10 \t\n2  \n1\t ",
			expected: "1\t \n2  \n10 \t",
		},
		{
			name:     "negative numbers with trailing blanks",
			input:    "-5  \n-10 \n-1",
			expected: "-10 \n-5  \n-1",
		},
		{
			name:     "decimal with trailing blanks",
			input:    "1.5 \n2.1\t\n1.2  ",
			expected: "1.2  \n1.5 \n2.1\t",
		},
		{
			name:     "non-numeric with trailing blanks",
			input:    "abc  \n5 \nxyz\t",
			expected: "5 \nabc  \nxyz\t",
		},
		{
			name:     "preserve original formatting",
			input:    "  10  \n 2 \n 1",
			expected: " 1\n 2 \n  10  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByNumeric(tt.input, trimBlanks, domain.SortOptions{})
			require.Equal(t, tt.expected, result)
		})
	}
}
