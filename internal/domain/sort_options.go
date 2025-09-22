package domain

type SortOptions struct {
	// Тип сортировки
	Field        int  // N
	Key          bool // flag -k N
	Numeric      bool // flag -n
	Month        bool // flag -M
	HumanNumeric bool // falg -h

	// Модификаторы
	Reverse      bool // flag -r
	IgnoreBlanks bool // flag -b
	Unique       bool // flag -u

	// Анализ
	Check bool // flag -c
}
