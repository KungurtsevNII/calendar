## Переменные
MAIN=./cmd/main.go
BIN_PATH=./bin
BINARY_NAME=calendar-service

## Команда форматирует код.
.PHONY: fmt
fmt:
	gofmt -w -s .

## Команда создаст каталог vendor в корне нашего проекта, содержащий исходный код всех зависимостей.
.PHONY: vendor
vendor:
	go mod vendor

## Команда собирает бинарник с использование vendor. Бинарки кладет в ./bin.
.PHONY: build
build: vendor
	GO111MODULE=on go build -mod=vendor -o=$(BIN_PATH)/$(BINARY_NAME) $(MAIN)

## Команда удаляет временные файлы, кэш и т.д.
.PHONY: clean
clean:
	rm -fr $(BIN_PATH)


