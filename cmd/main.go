package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unix_sort_lite/internal/domain"
	"unix_sort_lite/internal/usecase"

	"github.com/spf13/pflag"
)

func main() {
	// flags init
	field := pflag.IntP("key", "k", 0, "number of field")
	numeric := pflag.BoolP("numeric", "n", false, "numeric sort")
	month := pflag.BoolP("month-sort", "M", false, "month sort")
	humanNumeric := pflag.BoolP("human-numeric-sort", "h", false, "human numeric sort")
	reverse := pflag.BoolP("reverse", "r", false, "reverse")
	blanks := pflag.BoolP("ignore-trailing-blanks", "b", false, "ignore blanks")
	unique := pflag.BoolP("unique", "u", false, "unique")
	check := pflag.BoolP("check", "c", false, "check")

	pflag.Parse()

	opts := domain.SortOptions{
		Field:        *field,
		Numeric:      *numeric,
		Month:        *month,
		HumanNumeric: *humanNumeric,
		Reverse:      *reverse,
		IgnoreBlanks: *blanks,
		Unique:       *unique,
		Check:        *check,
	}
	pflag.Visit(func(f *pflag.Flag) {
		if f.Name == "key" {
			opts.Key = true
		}
	})

	args := pflag.Args()

	var input string
	if len(args) == 0 {
		// Читаем из stdin если файлы не указаны
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("failed to read from stdin:", err)
		}
		input = string(b)
	} else {
		// Читаем из файла
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("failed to open '%s': %v", args[0], err)
		}
		defer file.Close() //nolint:errcheck

		b, err := io.ReadAll(file)
		if err != nil {
			log.Fatal("failed to read file:", err)
		}
		input = string(b)
	}

	result, err := usecase.Sort(input, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.Check {
		if input != result {
			fmt.Println(domain.ErrWrongOrder)
			os.Exit(1)
		}
		return
	}

	fmt.Println(result)
}
