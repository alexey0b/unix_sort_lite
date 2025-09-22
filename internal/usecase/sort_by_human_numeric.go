package usecase

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unix_sort_lite/internal/domain"
)

var (
	numSISuffixRegex = regexp.MustCompile(`(?i)^\s*([-\+]?\d+\.?\d*)(K|M|G|T|P|E|Z|Y|R|Q)?\s*$`)
	SISuffixOrder    = map[string]int{
		"k": 1, "m": 2, "g": 3, "t": 4, "p": 5,
		"e": 6, "z": 7, "y": 8, "r": 9, "q": 10,
	}
)

// SortByHumanNumeric сортирует строки, учитывая SI суффиксы и поддерживает работу с вещественными числами.
func SortByHumanNumeric(s string, modify func(string) string, opts domain.SortOptions) string {
	rows := strings.Split(s, "\n")
	sort.SliceStable(rows, func(i, j int) bool {
		return compareHumanNumericStrings(modify(rows[i]), modify(rows[j]))
	})
	return strings.Join(rows, "\n")
}

// compareHumanNumericStrings сравнивает две строки по правилам human-readable сортировки.
// Реализует сложную логику сортировки согласно документации Unix sort -h:
// 1. По знаку числа (отрицательные < ноль < положительные)
// 2. По SI суффиксу (пустой < K < M < G...)
// 3. По числовому значению
//
// Примеры правильного порядка:
//
//	-2M < -1M < -2K < -1K < -2 < -1 < 0 < 1 < 2 < 1K < 2K < 1M < 2M < abc
func compareHumanNumericStrings(iStr, jStr string) bool {
	iHumNumMatch := numSISuffixRegex.FindStringSubmatch(iStr)
	jHumNumMatch := numSISuffixRegex.FindStringSubmatch(jStr)

	switch {
	case len(iHumNumMatch) == 3 && len(jHumNumMatch) == 3:
		iNum, _ := strconv.ParseFloat(iHumNumMatch[1], 64)
		jNum, _ := strconv.ParseFloat(jHumNumMatch[1], 64)

		if getSign(iNum) != getSign(jNum) {
			return getSign(iNum) < getSign(jNum)
		}

		iSuffix := getSuffixOrder(iHumNumMatch[2])
		jSuffix := getSuffixOrder(jHumNumMatch[2])
		if iSuffix != jSuffix {
			// Специальная обработка нулей: для них порядок суффиксов обычный
			if iNum == 0 || jNum == 0 {
				return iSuffix < jSuffix
			}
			// Это обращает порядок суффиксов для отрицательных чисел: -1M < -1K
			return iSuffix*getSign(iNum) < jSuffix*getSign(jNum)
		}

		return iNum < jNum
	case len(iHumNumMatch) == 3 && len(jHumNumMatch) != 3:
		return true
	case len(iHumNumMatch) != 3 && len(jHumNumMatch) == 3:
		return false
	default:
		return iStr < jStr
	}
}

// getSign возвращает знак числа для правильной сортировки.
// Используется для определения порядка: отрицательные < ноль < положительные
func getSign(num float64) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 0
	}
	return 1
}

// getSuffixOrder возвращает числовой порядок SI суффикса для сортировки.
// Преобразует суффикс в нижний регистр для регистронезависимого сравнения.
func getSuffixOrder(suffix string) int {
	return SISuffixOrder[strings.ToLower(suffix)]
}
