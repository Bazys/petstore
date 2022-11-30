#!make
ifneq (,$(wildcard ./.env))
    include ./.env
    export $(shell sed 's/=.*//' ./.env)
endif

PROJECT=petstore
HOME_DIR?=$(CURDIR)
GOLANGCI_LINT=$(shell go env GOPATH)/bin/golangci-lint
GOFUMPT=$(shell go env GOPATH)/bin/gofumpt

.PHONY: test
test:
	go test -v -race -cover -covermode=atomic ./internal/...

run: gomod fmt swag
	@go run ${HOME_DIR}/cmd/main.go

hooks:
	@git config core.hooksPath ./.githooks

lint: fmt swag
	${GOLANGCI_LINT} run --timeout 90s --exclude-use-default ./...

fmt:
	${GOFUMPT} -l -w -extra .

gomod:
	@go mod tidy

install:
	@echo "Installing dependencies..."
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go install github.com/google/gnostic@latest
	@go install github.com/googleapis/gnostic-grpc@latest
	@go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

goinit:
	@go mod init ${PROJECT}

gen-proto:
	@echo "Generating proto file from yaml..."
	@mkdir -p proto
	@gnostic --grpc-out=./proto ./api/${PROJECT}.json
# что бы можно было импортировать в другие пакеты
	@awk 'NR==2{print; print "option go_package = \"/pb\";"} NR!=2' proto/${PROJECT}.proto > tmp && mv -f tmp proto/${PROJECT}.proto

gen-pb:
	@echo "Generating pb,grpc & gw file from proto..."
	@mkdir -p pkg/pb
	@protoc -Iproto/ -Ivendor.proto/ \
		--govalidators_out ./pkg/ \
		--go_opt paths=source_relative \
		--go_out ./pkg/pb  \
		--go-grpc_opt paths=source_relative \
		--go-grpc_out ./pkg/pb  \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt generate_unbound_methods=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_out ./pkg/pb \
		${PROJECT}.proto
# в виде отдельного пакета
	@cd pkg/pb && go mod init ${PROJECT}/pkg/pb && go mod tidy && cd ../..

gen: vendor-grpc gen-proto gen-pb

vendor-grpc:
	@mkdir -p vendor.proto/google/api
	@echo "Installing google/api/annotations.proto..."
	@curl -o vendor.proto/google/api/annotations.proto -L'#' "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto"
	@echo "Installing google/api/http.proto..."
	@curl -o vendor.proto/google/api/http.proto -L'#' "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto"
	@echo "Installing google/api/field_behavior.proto..."
	@curl -o vendor.proto/google/api/field_behavior.proto -L'#' "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto"
	@echo "Installing google/api/httpbody.proto..."
	@curl -o vendor.proto/google/api/httpbody.proto -L'#' "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto"
	@echo "Installing github.com/mwitkow/go-proto-validators..."
	@mkdir -p vendor.proto/github.com/mwitkow/go-proto-validators
	@curl -o vendor.proto/github.com/mwitkow/go-proto-validators/validator.proto -L'#' "https://raw.githubusercontent.com/mwitkow/go-proto-validators/master/validator.proto"
