package usecase

import (
	"strings"
)

// Reverse переворачивает порядок строк в обратном направлении.
// Реализует функциональность флага -r (reverse) в Unix sort.
//
// Примеры:
//
//	"a\nb\nc" → "c\nb\na"
//	"1\n2\n3\n4" → "4\n3\n2\n1"
//	"single" → "single" (одна строка остается без изменений)
func Reverse(s string) string {
	rows := strings.Split(s, "\n")
	l, r := 0, len(rows)-1
	for l < r {
		rows[l], rows[r] = rows[r], rows[l]
		l++
		r--
	}
	return strings.Join(rows, "\n")
}
