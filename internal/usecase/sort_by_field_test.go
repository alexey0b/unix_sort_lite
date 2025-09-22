package usecase

import (
	"testing"
	"unix_sort_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestSortByField(t *testing.T) {
	identity := func(s string) string { return s }

	tests := []struct {
		name     string
		input    string
		opts     domain.SortOptions
		expected string
	}{
		{
			name:     "sort by field 2 lexicographic",
			input:    "d Feb\nc apr\na\nf",
			opts:     domain.SortOptions{Field: 2},
			expected: "a\nf\nc apr\nd Feb",
		},
		{
			name:     "sort by field 1",
			input:    "zebra\napple\nbanana",
			opts:     domain.SortOptions{Field: 1},
			expected: "apple\nbanana\nzebra",
		},
		{
			name:     "sort by field 3",
			input:    "a b c\nx y z\nm n o",
			opts:     domain.SortOptions{Field: 3},
			expected: "a b c\nm n o\nx y z",
		},
		{
			name:     "mixed case sorting",
			input:    "a Apple\nb banana\nc Cherry",
			opts:     domain.SortOptions{Field: 2},
			expected: "a Apple\nb banana\nc Cherry",
		},
		{
			name:     "lines without field",
			input:    "single\na b\nc d e",
			opts:     domain.SortOptions{Field: 2},
			expected: "single\na b\nc d e",
		},
		{
			name:     "numeric field sort",
			input:    "a 10\nb 2\nc 1",
			opts:     domain.SortOptions{Field: 2, Numeric: true},
			expected: "c 1\nb 2\na 10",
		},
		{
			name:     "month field sort",
			input:    "event Feb\nparty Jan\nmeeting Mar",
			opts:     domain.SortOptions{Field: 2, Month: true},
			expected: "party Jan\nevent Feb\nmeeting Mar",
		},
		{
			name:     "human numeric field sort",
			input:    "file 1M\ndata 2K\nlog 500",
			opts:     domain.SortOptions{Field: 2, HumanNumeric: true},
			expected: "log 500\ndata 2K\nfile 1M",
		},
		{
			name:     "empty input",
			input:    "",
			opts:     domain.SortOptions{Field: 1},
			expected: "",
		},
		{
			name:     "single line",
			input:    "hello world",
			opts:     domain.SortOptions{Field: 2},
			expected: "hello world",
		},
		{
			name:     "field beyond available",
			input:    "a\nb c\nd e f",
			opts:     domain.SortOptions{Field: 5},
			expected: "a\nb c\nd e f",
		},
		{
			name:     "whitespace handling",
			input:    "a   b\nc d\ne    f   g",
			opts:     domain.SortOptions{Field: 2},
			expected: "a   b\nc d\ne    f   g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByField(tt.input, identity, tt.opts)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSortByFieldWithModify(t *testing.T) {
	trimBlanks := func(s string) string {
		return ignoreTrailingBlanks(s)
	}

	tests := []struct {
		name     string
		input    string
		opts     domain.SortOptions
		expected string
	}{
		{
			name:     "ignore trailing blanks in field",
			input:    "a apple  \nb banana\nc cherry ",
			opts:     domain.SortOptions{Field: 2},
			expected: "a apple  \nb banana\nc cherry ",
		},
		{
			name:     "trailing blanks affect sorting",
			input:    "a b  \nc a\nd b",
			opts:     domain.SortOptions{Field: 2},
			expected: "c a\na b  \nd b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByField(tt.input, trimBlanks, tt.opts)
			require.Equal(t, tt.expected, result)
		})
	}
}
