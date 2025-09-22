package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortDefault(t *testing.T) {
	identity := func(s string) string { return s }

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic alphabetical sort",
			input:    "c\nb\na",
			expected: "a\nb\nc",
		},
		{
			name:     "mixed case sorting",
			input:    "Zebra\napple\nBanana",
			expected: "Banana\nZebra\napple",
		},
		{
			name:     "numbers as strings",
			input:    "10\n2\n1",
			expected: "1\n10\n2",
		},
		{
			name:     "special characters",
			input:    "!test\n@test\n#test",
			expected: "!test\n#test\n@test",
		},
		{
			name:     "empty lines",
			input:    "c\n\nb\na",
			expected: "\na\nb\nc",
		},
		{
			name:     "whitespace lines",
			input:    "c\n  \nb\na",
			expected: "  \na\nb\nc",
		},
		{
			name:     "identical lines",
			input:    "same\nsame\nsame",
			expected: "same\nsame\nsame",
		},
		{
			name:     "single line",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "unicode characters",
			input:    "ñ\nz\na\nä",
			expected: "a\nz\nä\nñ",
		},
		{
			name:     "long strings",
			input:    "verylongstring\nshort\nmediumlength",
			expected: "mediumlength\nshort\nverylongstring",
		},
		{
			name:     "strings with spaces",
			input:    "hello world\nhello\nhello there",
			expected: "hello\nhello there\nhello world",
		},
		{
			name:     "mixed alphanumeric",
			input:    "file10\nfile2\nfile1",
			expected: "file1\nfile10\nfile2",
		},
		{
			name:     "punctuation sorting",
			input:    "test.\ntest,\ntest!",
			expected: "test!\ntest,\ntest.",
		},
		{
			name:     "case sensitivity order",
			input:    "A\na\nB\nb",
			expected: "A\nB\na\nb",
		},
		{
			name:     "leading spaces",
			input:    " apple\napple\n  apple",
			expected: "  apple\n apple\napple",
		},
		{
			name:     "tabs vs spaces",
			input:    "\tapple\n apple\napple",
			expected: "\tapple\n apple\napple",
		},
		{
			name:     "numeric strings with leading zeros",
			input:    "001\n01\n1",
			expected: "001\n01\n1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortDefault(tt.input, identity)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSortDefaultWithModify(t *testing.T) {
	trimBlanks := func(s string) string {
		return ignoreTrailingBlanks(s)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trailing spaces ignored",
			input:    "apple  \nbanana \ncherry",
			expected: "apple  \nbanana \ncherry",
		},
		{
			name:     "trailing tabs ignored",
			input:    "apple\t\nbanana\t\t\ncherry",
			expected: "apple\t\nbanana\t\t\ncherry",
		},
		{
			name:     "mixed trailing blanks",
			input:    "apple \t\nbanana  \ncherry\t ",
			expected: "apple \t\nbanana  \ncherry\t ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortDefault(tt.input, trimBlanks)
			require.Equal(t, tt.expected, result)
		})
	}
}
