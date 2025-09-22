package usecase

import (
	"sort"
	"strings"
)

// SortDefault выполняет лексикографическую сортировку (по умолчанию в Unix sort).
// Сортирует строки по ASCII/Unicode значениям символов, символ за символом.
// Поддерживает модификацию строк перед сравнением (например, для флага -b).
//
// Примеры:
//
//	"c\nb\na" → "a\nb\nc"
//	"Zebra\napple\nBanana" → "Banana\nZebra\napple" (заглавные первыми)
//	"10\n2\n1" → "1\n10\n2" (лексикографически, не числово)
//	"apple  \nbanana \ncherry" с -b → сравнение без trailing пробелов
func SortDefault(s string, modify func(string) string) string {
	rows := strings.Split(s, "\n")
	sort.SliceStable(rows, func(i, j int) bool {
		return modify(rows[i]) < modify(rows[j])
	})
	return strings.Join(rows, "\n")
}
