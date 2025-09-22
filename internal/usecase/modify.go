package usecase

import (
	"strings"
)

// trailingBlanks определяет символы, которые считаются trailing blanks
// Включает пробел и табуляцию в соответствии с поведением Unix sort -b
const trailingBlanks = " \t"

// ignoreTrailingBlanks удаляет trailing пробелы и табуляции из строки.
// Используется для реализации флага -b (ignore trailing blanks).
// Важно: функция удаляет только trailing blanks, leading пробелы сохраняются.
//
// Примеры:
//
//	"hello  " → "hello"
//	"  hello  " → "  hello"
//	"hello\t\t" → "hello"
func ignoreTrailingBlanks(s string) string {
	return strings.TrimRight(s, trailingBlanks)
}
