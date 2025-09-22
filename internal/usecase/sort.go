package usecase

import (
	"unix_sort_lite/internal/domain"
)

// Sort выполняет сортировку входных данных согласно переданным опциям.
// Поддерживает различные типы сортировки и модификаторы в стиле Unix sort.
func Sort(input string, opts domain.SortOptions) (string, error) {
	var (
		sortTypes int
		result    string
		modify    func(string) string
	)

	// Валидация: флаг -k требует корректный номер поля
	if opts.Key && opts.Field < 1 {
		return "", domain.ErrInvalideField
	}

	// Подсчет взаимоисключающих типов сортировки
	if opts.Numeric {
		sortTypes++
	}
	if opts.Month {
		sortTypes++
	}
	if opts.HumanNumeric {
		sortTypes++
	}
	// Проверка конфликтующих флагов, например, -nM
	if sortTypes > 1 {
		return "", domain.ErrConflictOpts
	}

	if opts.IgnoreBlanks {
		modify = ignoreTrailingBlanks
	} else {
		modify = func(s string) string { return s } // Identity функция
	}

	switch {
	case opts.Key:
		// Сортировка по N-му полю -k флаг
		result = SortByField(input, modify, opts)
	case opts.Numeric:
		// Числовая сортировка -n флаг
		result = SortByNumeric(input, modify, opts)
	case opts.Month:
		// Сортировка по месяцам -M флаг
		result = SortByMonth(input, modify, opts)
	case opts.HumanNumeric:
		// Human-readable сортировка -h флаг
		result = SortByHumanNumeric(input, modify, opts)
	default:
		// Лексикографическая сортировка по умолчанию
		result = SortDefault(input, modify)
	}

	if opts.Reverse {
		result = Reverse(result)
	}
	if opts.Unique {
		result = Unique(result, opts.Field)
	}

	return result, nil
}
