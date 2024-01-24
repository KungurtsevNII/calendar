## Переменные
## Входная точка(папка) для компиляции(старта) приложения.
MAIN=$(PWD)/cmd/*.go
## Папка с бинарями.
BIN_PATH=$(PWD)/bin
## Имя бинарники сервиса.
BINARY_NAME=calendar-service
## Папка tmp, в которую скачиваются архивированные бинарники.
TMP=$(BIN_PATH)/tmp

## Переменные для proto.
## Версии proto генераторов.
PROTOC_GEN_GO_VERSION=v1.5.3
PROTOC_GEN_GO_GRPC_VERSION=v1.3.0
PROTOC_VERSION=23.3
## Полные имена бинарников для proto генерации.
PROTOC_BIN=$(BIN_PATH)/protoc
PROTOC_GEN_GO_BIN=$(BIN_PATH)/protoc-gen-go
PROTOC_GEN_GO_GRPC_BIN=$(BIN_PATH)/protoc-gen-go-grpc
## Папка, в которую складываются все сгенерированные файлы.
PROTO_OUT=$(PWD)/pkg/pb
## Репа от куда качаем protoc.
PB_REL=https://github.com/protocolbuffers/protobuf/releases
## Определяет OS и архитектуру процессора для скачивания protoc бинаря.
PROTOC_OS_ARC :=
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	PROTOC_OS_ARC :=linux
endif
ifeq ($(UNAME_S),Darwin)
	PROTOC_OS_ARC :=osx
endif
UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
	PROTOC_OS_ARC :=$(PROTOC_OS_ARC)-x86_64
endif
ifneq ($(filter %86,$(UNAME_M)),)
	PROTOC_OS_ARC :=$(PROTOC_OS_ARC)-x86_32
endif
ifneq ($(filter arm%,$(UNAME_M)),)
	PROTOC_OS_ARC :=$(PROTOC_OS_ARC)-aarch_64
endif

## Команда форматирует код.
.PHONY: fmt
fmt:
	gofmt -w -s .

## Команда устанавливает бинарники, которые требуются для работы с proto.
.PHONY: install-proto-generator
install-proto-generator:
	mkdir -p $(TMP)
	{ \
	if ! [ -f ${TMP}/protoc-${PROTOC_VERSION}.zip ]; then \
		curl -#fLo ${TMP}/protoc-${PROTOC_VERSION}.zip "${PB_REL}/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_OS_ARC}.zip"; \
	fi; \
	}
	unzip -o -u ${TMP}/protoc-${PROTOC_VERSION}.zip -d ${TMP}/protoc-${PROTOC_VERSION}
	cp -rf ${TMP}/protoc-${PROTOC_VERSION}/bin/protoc $(BIN_PATH)
	##cp -rf ${TMP}/protoc-${PROTOC_VERSION}/include/* $(BIN_PATH)/include/

	GO111MODULE=on GOBIN=$(BIN_PATH) go install github.com/golang/protobuf/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	GO111MODULE=on GOBIN=$(BIN_PATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

## Команда компиоирует proto и генерирует go код.
.PHONY: generate-proto
generate-proto:install-proto-generator
	mkdir -p $(PROTO_OUT)
	$(PROTOC_BIN) $(PWD)/api/calendar.proto \
		--proto_path $(PWD)/api \
		--go_out=$(PROTO_OUT) \
		--go_opt=paths=source_relative \
		--plugin=protoc-gen-go=$(PROTOC_GEN_GO_BIN) \
		--plugin=protoc-gen-go-grpc=$(PROTOC_GEN_GO_GRPC_BIN) \
		--go-grpc_out=$(PROTO_OUT) \
		--go-grpc_opt=paths=source_relative

## Команда создаст каталог vendor в корне нашего проекта, содержащий исходный код всех зависимостей.
.PHONY: vendor
vendor:
	go mod vendor

## Команда собирает бинарник с использование vendor. Бинарки кладет в ./bin.
.PHONY: build
build: generate-proto vendor
	mkdir -p $(BIN_PATH)
	GO111MODULE=on go build -mod=vendor -o=$(BIN_PATH)/$(BINARY_NAME) $(MAIN)

## Команда запускает скомпилированный бинарник. Передает cli параметр, который указывает путь к локальным конфигам.
.PHONY: run-local
run-local: build
	$(BIN_PATH)/$(BINARY_NAME) -config-path=./config/config_local.yaml

## Команда удаляет временные файлы, кэш и т.д.
.PHONY: clean
clean:
	rm -fr $(BIN_PATH)
	rm -fr $(TMP)
	rm -fr $(PWD)/vendor
	rm -fr $(PROTO_OUT)


