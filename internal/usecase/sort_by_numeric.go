package usecase

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unix_sort_lite/internal/domain"
)

var numRegex = regexp.MustCompile(`^\s*([-+]?)(\d+\.?\d*)\s*$`)

// SortByNumeric выполняет числовую сортировку (флаг -n в Unix sort).
// Распознает чистые числа (целые и десятичные) с опциональными знаками и сортирует их
// в числовом порядке. Строки, не являющиеся числами, сортируются лексикографически
// и идут после чисел.
//
// Примеры:
//
//	"10\n2\n1" → "1\n2\n10" (числовая сортировка, не лексикографическая)
//	"5\nabc\n-3\nxyz" → "-3\n5\nabc\nxyz" (числа первыми, потом не-числа)
//	"-5.5\n2.1\n0" → "-5.5\n0\n2.1" (поддержка отрицательных и десятичных)
func SortByNumeric(s string, modify func(string) string, opts domain.SortOptions) string {
	rows := strings.Split(s, "\n")
	sort.SliceStable(rows, func(i, j int) bool {
		return compareNumericStrings(modify(rows[i]), modify(rows[j]))
	})
	return strings.Join(rows, "\n")
}

// compareNumericStrings сравнивает две строки по правилам числовой сортировки.
//
// Примеры правильного порядка:
//
//	-10 < -5 < -1 < 0 < 1 < 2 < 10 < abc < xyz
func compareNumericStrings(iStr, jStr string) bool {
	iMatch, jMatch := numRegex.FindStringSubmatch(iStr), numRegex.FindStringSubmatch(jStr)

	switch {
	case len(iMatch) == 3 && len(jMatch) == 3:
		iNum, _ := strconv.ParseFloat(iMatch[2], 64)
		jNum, _ := strconv.ParseFloat(jMatch[2], 64)

		if getSignAsNum(iMatch[1]) != getSignAsNum(jMatch[1]) {
			return getSignAsNum(iMatch[1]) < getSignAsNum(jMatch[1])
		}

		// Знаки одинаковые, сравниваем полные числа (число * знак)
		return getSignAsNum(iMatch[1])*iNum < getSignAsNum(jMatch[1])*jNum

	case len(iMatch) == 3 && len(jMatch) != 3:
		return true
	case len(iMatch) != 3 && len(jMatch) == 3:
		return false
	default:
		return iStr < jStr
	}
}

// getSignAsNum преобразует строковый знак в числовое значение для математических операций.
// Используется для правильного сравнения положительных и отрицательных чисел.
//
// Примеры:
//
//	getSignAsNum("-") → -1.0 (для чисел вида "-123")
//	getSignAsNum("") → 1.0 (для чисел вида "123")
//	getSignAsNum("+") → 1.0 (для чисел вида "+123")
func getSignAsNum(sign string) float64 {
	if sign == "-" {
		return -1
	}
	return 1
}
