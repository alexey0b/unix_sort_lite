package usecase

import (
	"regexp"
	"sort"
	"strings"
	"unix_sort_lite/internal/domain"
)

var (
	monthRegex = regexp.MustCompile(`(?i)^\s*(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)\s*$`)

	// monthOrder определяет порядок месяцев в году для сортировки
	// Соответствует календарному порядку: январь (1) → декабрь (12)
	monthOrder = map[string]int{
		"jan": 1, "feb": 2, "mar": 3, "apr": 4,
		"may": 5, "jun": 6, "jul": 7, "aug": 8,
		"sep": 9, "oct": 10, "nov": 11, "dec": 12,
	}
)

// SortByMonth выполняет сортировку по месяцам (флаг -M в Unix sort).
// Распознает сокращенные названия месяцев (Jan, Feb, Mar, ...) и сортирует их
// в календарном порядке. Строки, не содержащие месяцы, сортируются
// лексикографически и идут первыми.
//
// Примеры:
//
//	"Feb\nJan\nMar" → "Jan\nFeb\nMar"
//	"abc\nFeb\nxyz\nJan" → "abc\nxyz\nJan\nFeb" (не-месяцы первыми)
func SortByMonth(s string, modify func(string) string, opts domain.SortOptions) string {
	rows := strings.Split(s, "\n")
	sort.SliceStable(rows, func(i, j int) bool {
		return compareMonthStrings(modify(rows[i]), modify(rows[j]))
	})
	return strings.Join(rows, "\n")
}

// compareMonthStrings сравнивает две строки по правилам месячной сортировки.
// Примеры порядка:
//
//	"abc" < "xyz" < "Jan" < "Feb" < "Mar" < ... < "Dec"
func compareMonthStrings(iStr, jStr string) bool {
	iMatch, jMatch := monthRegex.FindStringSubmatch(strings.ToLower(iStr)), monthRegex.FindStringSubmatch(strings.ToLower(jStr))

	switch {
	case len(iMatch) == 2 && len(jMatch) == 2:
		return monthOrder[iMatch[1]] < monthOrder[jMatch[1]]
	case len(iMatch) == 2 && len(jMatch) != 2:
		return false
	case len(iMatch) != 2 && len(jMatch) == 2:
		return true
	default:
		return iStr < jStr
	}
}
