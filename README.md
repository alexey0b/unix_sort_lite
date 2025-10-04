[![Unix Sort Lite CI](https://github.com/alexey0b/unix_sort_lite/actions/workflows/ci.yaml/badge.svg)](https://github.com/alexey0b/unix_sort_lite/actions/workflows/ci.yaml)
[![Coverage Status](https://coveralls.io/repos/github/alexey0b/unix_sort_lite/badge.svg?branch=main)](https://coveralls.io/github/alexey0b/unix_sort_lite?branch=main)

# 🧘🏼‍♀️ Unix Sort Lite

Легковесная реализация утилиты `sort` на Go с поддержкой основных флагов Unix sort.

---

## ✅ Поддерживаемые флаги

| Флаг                           | Описание                 | Пример                                           |
| ------------------------------ | ------------------------ | ------------------------------------------------ |
| `-n, --numeric`                | Числовая сортировка      | `echo -e "10\n2\n1" \| ./unix_sort_lite -n`      |
| `-r, --reverse`                | Обратная сортировка      | `echo -e "a\nb\nc" \| ./unix_sort_lite -r`       |
| `-k, --key N`                  | Сортировка по полю N     | `echo -e "b 2\na 1" \| ./unix_sort_lite -k 2`    |
| `-M, --month-sort`             | Сортировка по месяцам    | `echo -e "Mar\nJan\nFeb" \| ./unix_sort_lite -M` |
| `-h, --human-numeric-sort`     | Человеко-читаемые числа  | `echo -e "1K\n2M\n3G" \| ./unix_sort_lite -h`    |
| `-u, --unique`                 | Только уникальные строки | `echo -e "a\na\nb" \| ./unix_sort_lite -u`       |
| `-b, --ignore-trailing-blanks` | Игнорировать пробелы     | `echo -e " a\nb " \| ./unix_sort_lite -b`        |
| `-c, --check`                  | Проверить сортировку     | `echo -e "a\nb\nc" \| ./unix_sort_lite -c`       |

---

## ▶️ Использование

- **Клонируйте репозиторий:**

```bash
git clone https://github.com/alexey0b/unix_sort_lite
```

- **Посмотрите все доступные команды**

```bash
make help
```

---

## Примеры использования утилиты на файле `input.txt`

### Стандартная сортировка

```bash
make sort INPUT_FILE='example/words.txt'
# Output:
alert
go
package
test
```

### Числовая сортировка

```bash
make sort FLAGS='-n' INPUT_FILE='example/nums.txt'
# Output:
1
2
10
```

### Сортировка по месяцам

```bash
make sort FLAGS='-M' INPUT_FILE='example/months.txt'
# Output:
Jan
Feb
Mar
```

### Комбинированные флаги

```bash
make sort FLAGS='-nr' INPUT_FILE='example/nums.txt'
# Output:
10
2
1
```

---

## 🛠️ Технические ресурсы

### Требования

- Go 1.18+
- Unix/Linux/macOS
- golangci-lint для линтинга (опционально)

### Зависимости

- **[spf13/pflag](https://github.com/spf13/pflag)** - POSIX/GNU-style флаги

---
