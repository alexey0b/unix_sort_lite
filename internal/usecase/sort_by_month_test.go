package usecase

import (
	"testing"
	"unix_sort_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestSortByMonth(t *testing.T) {
	identity := func(s string) string { return s }

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic month sort",
			input:    "Feb\nJan\nMar",
			expected: "Jan\nFeb\nMar",
		},
		{
			name:     "all months in order",
			input:    "Dec\nNov\nOct\nSep\nAug\nJul\nJun\nMay\nApr\nMar\nFeb\nJan",
			expected: "Jan\nFeb\nMar\nApr\nMay\nJun\nJul\nAug\nSep\nOct\nNov\nDec",
		},
		{
			name:     "mixed case months",
			input:    "feb\nJAN\nMar",
			expected: "JAN\nfeb\nMar",
		},
		{
			name:     "months with non-month strings",
			input:    "abc\nFeb\nxyz\nJan",
			expected: "abc\nxyz\nJan\nFeb",
		},
		{
			name:     "duplicate months",
			input:    "Jan\nFeb\nJan\nMar",
			expected: "Jan\nJan\nFeb\nMar",
		},
		{
			name:     "months in text",
			input:    "event Feb\nparty Jan\nmeeting Mar",
			expected: "event Feb\nmeeting Mar\nparty Jan",
		},
		{
			name:     "partial month names",
			input:    "January\nFebruary\nMarch",
			expected: "February\nJanuary\nMarch",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "single month",
			input:    "Jan",
			expected: "Jan",
		},
		{
			name:     "no months",
			input:    "apple\nbanana\ncherry",
			expected: "apple\nbanana\ncherry",
		},
		{
			name:     "months at different positions",
			input:    "data-Feb-2024\nlog-Jan-2024\nreport-Mar-2024",
			expected: "data-Feb-2024\nlog-Jan-2024\nreport-Mar-2024",
		},
		{
			name:     "mixed months and numbers",
			input:    "123\nFeb\n456\nJan",
			expected: "123\n456\nJan\nFeb",
		},
		{
			name:     "summer months",
			input:    "Aug\nJul\nJun",
			expected: "Jun\nJul\nAug",
		},
		{
			name:     "winter months",
			input:    "Feb\nDec\nJan",
			expected: "Jan\nFeb\nDec",
		},
		{
			name:     "spring months",
			input:    "May\nMar\nApr",
			expected: "Mar\nApr\nMay",
		},
		{
			name:     "fall months",
			input:    "Nov\nSep\nOct",
			expected: "Sep\nOct\nNov",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByMonth(tt.input, identity, domain.SortOptions{})
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCompareMonthStrings(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected bool
	}{
		{
			name:     "Jan befor Feb",
			a:        "Jan",
			b:        "Feb",
			expected: true,
		},
		{
			name:     "Dec after Jan",
			a:        "Dec",
			b:        "Jan",
			expected: false,
		},
		{
			name:     "non-month after month",
			a:        "abc",
			b:        "Jan",
			expected: true,
		},
		{
			name:     "month befor non-month",
			a:        "Jan",
			b:        "abc",
			expected: false,
		},
		{
			name:     "non-month < non-month",
			a:        "abc",
			b:        "xyz",
			expected: true,
		},
		{
			name:     "case insensitive",
			a:        "jan",
			b:        "FEB",
			expected: true,
		},
		{
			name:     "same month",
			a:        "Jan",
			b:        "Jan",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := compareMonthStrings(tt.a, tt.b)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSortByMonthWithModify(t *testing.T) {
	trimBlanks := func(s string) string {
		return ignoreTrailingBlanks(s)
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "months with trailing spaces",
			input:    "Feb  \nJan \nMar",
			expected: "Jan \nFeb  \nMar",
		},
		{
			name:     "months with trailing tabs",
			input:    "Feb\t\nJan\t\t\nMar",
			expected: "Jan\t\t\nFeb\t\nMar",
		},
		{
			name:     "mixed trailing blanks",
			input:    "Feb \t\nJan  \nMar\t ",
			expected: "Jan  \nFeb \t\nMar\t ",
		},
		{
			name:     "non-months with trailing blanks",
			input:    "abc  \nFeb \nxyz\t",
			expected: "abc  \nxyz\t\nFeb ",
		},
		{
			name:     "preserve original formatting",
			input:    "  Feb  \n Jan \n Mar",
			expected: " Jan \n  Feb  \n Mar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SortByMonth(tt.input, trimBlanks, domain.SortOptions{})
			require.Equal(t, tt.expected, result)
		})
	}
}
