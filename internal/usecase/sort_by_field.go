package usecase

import (
	"sort"
	"strings"
	"unix_sort_lite/internal/domain"
)

// rowData представляет строку данных с разделенными полями и оригинальным содержимым.
type rowData struct {
	fields   []string
	original string
}

// SortByField выполняет сортировку по указанному полю (флаг -k в Unix sort).
// Поддерживает различные типы интерпретации поля: числовая, месячная, human-readable.
// Строки без достаточного количества полей сортируются по количеству полей.

// Примеры:
//
//	"apple red\nbanana yellow" с полем 2 → сортировка по "red", "yellow"
//	"a\nb c\nd e f" с полем 3 → строки без поля 3 идут первыми
func SortByField(s string, modify func(string) string, opts domain.SortOptions) string {
	N := opts.Field
	lines := strings.Split(s, "\n")

	// Создаем массив структур для хранения полей и оригинальных строк
	rows := make([]rowData, len(lines))
	for i, line := range lines {
		rows[i] = rowData{
			fields:   strings.Fields(line),
			original: line,
		}
	}

	sort.SliceStable(rows, func(i, j int) bool {
		// Обработка строк с недостаточным количеством полей
		if len(rows[i].fields) < N || len(rows[j].fields) < N {
			return len(rows[i].fields) < len(rows[j].fields)
		}

		iField := modify(rows[i].fields[N-1])
		jField := modify(rows[j].fields[N-1])

		switch {
		case opts.Numeric:
			// Числовое сравнение полей (флаг -n)
			return compareNumericStrings(iField, jField)
		case opts.Month:
			// Сравнение по месяцам (флаг -M)
			return compareMonthStrings(iField, jField)
		case opts.HumanNumeric:
			// Human-readable числовое сравнение (флаг -h)
			return compareHumanNumericStrings(iField, jField)
		default:
			// Лексикографическое сравнение (по умолчанию)
			// Используем ToLower для регистронезависимого сравнения
			return strings.ToLower(iField) < strings.ToLower(jField)
		}
	})

	resLines := make([]string, len(rows))
	for i, row := range rows {
		resLines[i] = row.original
	}

	return strings.Join(resLines, "\n")
}
