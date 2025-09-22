package usecase

import (
	"strings"
)

// Unique удаляет дубликаты строк (флаг -u в Unix sort).
// Может работать как с целыми строками, так и с определенными полями в строках.
// Сохраняет порядок первого вхождения каждой уникальной строки/поля.
//
// Примеры:
//
//	"apple\nbanana\napple\ncherry" с field=0 → "apple\nbanana\ncherry"
//	"apple red\nbanana yellow\napple green" с field=1 → "apple red\nbanana yellow" (уникальность по 1-му полю)
//	"apple red\nbanana yellow\ngrape red" с field=2 → "apple red\nbanana yellow" (уникальность по 2-му полю)
//	"apple\nbanana yellow" с field=3 → "apple\nbanana yellow" (строки без поля 3 считаются дубликатами)
func Unique(s string, field int) string {
	lines := strings.Split(s, "\n")
	// Словарь для отслеживания уже встреченных ключей (строк или полей)
	dict := make(map[string]bool)
	uniqueRows := make([]string, 0, len(lines))

	for _, line := range lines {
		var key string

		if field < 1 {
			key = line
		} else {
			// Уникальность по N-му полю (комбинация -uk N)
			// Разделяем строку на поля по пробелам/табуляциям
			fields := strings.Fields(line)

			if len(fields) < field {
				// Строка не содержит достаточно полей
				key = ""
			} else {
				key = fields[field-1]
			}
		}

		if !dict[key] {
			uniqueRows = append(uniqueRows, line)
			dict[key] = true
		}
	}

	return strings.Join(uniqueRows, "\n")
}
