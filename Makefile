# Переменные
BINARY_NAME=l1_go_tasks
MAIN_PATH=./main.go
GOLANGCI_LINT_VERSION=v2.4.0

# Цвета для вывода
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

.PHONY: help build run clean test lint fmt vet deps install-tools

# Команда по умолчанию
.DEFAULT_GOAL := help

## help: показать справку по командам
help:
	@echo "Доступные команды:"
	@echo ""
	@echo "  ${GREEN}build${NC}        - собрать приложение"
	@echo "  ${GREEN}run${NC}          - запустить приложение"
	@echo "  ${GREEN}clean${NC}        - очистить собранные файлы"
	@echo ""
	@echo "  ${YELLOW}test${NC}         - запустить тесты"
	@echo "  ${YELLOW}test-verbose${NC} - запустить тесты с подробным выводом"
	@echo "  ${YELLOW}test-cover${NC}   - запустить тесты с покрытием"
	@echo ""
	@echo "  ${YELLOW}lint${NC}         - запустить линтер"
	@echo "  ${YELLOW}lint-fix${NC}     - запустить линтер с автоисправлениями"
	@echo "  ${YELLOW}fmt${NC}          - отформатировать код"
	@echo "  ${YELLOW}vet${NC}          - проверить код с go vet"
	@echo ""
	@echo "  ${RED}deps${NC}         - скачать зависимости"
	@echo "  ${RED}install-tools${NC} - установить необходимые инструменты"
	@echo "  ${RED}mod-tidy${NC}     - очистить go.mod"

## build: собрать приложение
build:
	@echo "${GREEN}Сборка приложения...${NC}"
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "${GREEN}Приложение собрано: bin/$(BINARY_NAME)${NC}"

## run: запустить приложение
run:
	@echo "${GREEN}Запуск приложения...${NC}"
	go run $(MAIN_PATH)

## clean: очистить собранные файлы
clean:
	@echo "${YELLOW}Очистка...${NC}"
	rm -rf bin/
	go clean
	@echo "${GREEN}Очистка завершена${NC}"

## test: запустить тесты
test:
	@echo "${YELLOW}Запуск тестов...${NC}"
	go test ./...

## test-verbose: запустить тесты с подробным выводом
test-verbose:
	@echo "${YELLOW}Запуск тестов (подробно)...${NC}"
	go test -v ./...

## test-cover: запустить тесты с покрытием
test-cover:
	@echo "${YELLOW}Запуск тестов с покрытием...${NC}"
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "${GREEN}Отчет о покрытии: coverage.html${NC}"

## lint: запустить линтер
lint:
	@echo "${YELLOW}Запуск линтера...${NC}"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "${RED}golangci-lint не установлен. Запустите: make install-tools${NC}"; \
		exit 1; \
	fi

## lint-fix: запустить линтер с автоисправлениями
lint-fix:
	@echo "${YELLOW}Запуск линтера с автоисправлениями...${NC}"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run --fix; \
	else \
		echo "${RED}golangci-lint не установлен. Запустите: make install-tools${NC}"; \
		exit 1; \
	fi

## fmt: отформатировать код
fmt:
	@echo "${YELLOW}Форматирование кода...${NC}"
	go fmt ./...
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	else \
		echo "${YELLOW}goimports не установлен, используется только go fmt${NC}"; \
	fi

## vet: проверить код с go vet
vet:
	@echo "${YELLOW}Проверка кода с go vet...${NC}"
	go vet ./...

## deps: скачать зависимости
deps:
	@echo "${YELLOW}Скачивание зависимостей...${NC}"
	go mod download
	@echo "${GREEN}Зависимости скачаны${NC}"

## mod-tidy: очистить go.mod
mod-tidy:
	@echo "${YELLOW}Очистка go.mod...${NC}"
	go mod tidy
	@echo "${GREEN}go.mod очищен${NC}"

## install-tools: установить необходимые инструменты
install-tools:
	@echo "${YELLOW}Установка инструментов...${NC}"
	@echo "Установка golangci-lint..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin latest; \
	else \
		echo "golangci-lint уже установлен"; \
	fi
	@echo "Установка goimports..."
	@if ! command -v goimports >/dev/null 2>&1; then \
		go install golang.org/x/tools/cmd/goimports@latest; \
	else \
		echo "goimports уже установлен"; \
	fi
	@echo "${GREEN}Инструменты установлены${NC}"

## check: полная проверка кода (fmt + vet + lint + test)
check: fmt vet lint test
	@echo "${GREEN}Полная проверка завершена${NC}"

## ci: команды для CI/CD
ci: deps check
	@echo "${GREEN}CI проверки завершены${NC}"
