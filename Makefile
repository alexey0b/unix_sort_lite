# Сборка утилиты
build:
	@go build -o unix_sort_lite cmd/main.go

# Запуск утилиты
sort:build
ifdef FLAGS
	@./unix_sort_lite $(FLAGS) $(INPUT_FILE)
else
	@./unix_sort_lite $(INPUT_FILE)
endif

# Тестирование
test:
	@go test -v ./internal/usecase

test-cover:
	@go test -v -covermode=atomic -coverprofile=coverage.out ./internal/usecase

# Качество кода
fmt:
	@go fmt ./...

vet:
	@go vet ./...

lint:
	@golangci-lint run

# Справка
help:
	@echo "Unix Sort Lite - Available commands:"
	@echo ""
	@echo "Build command:"
	@echo "  build         - Build binary for current OS"
	@echo ""
	@echo "Run commands:"
	@echo "  sort          - Sort file (FLAGS='-n' INPUT_FILE='input.txt')"
	@echo "                  If INPUT_FILE empty, reads from stdin"
	@echo "                  If FLAGS empty, uses default sorting"
	@echo ""
	@echo "Test commands:"
	@echo "  test          - Run all tests"
	@echo "  test-cover    - Run tests with coverage"
	@echo ""
	@echo "Code quality:"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run golangci-lint"

.PHONY: build sort test test-cover help fmt vet lint
