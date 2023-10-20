entrypoint=./cmd/sylphy
tmp_bin=./tmp/bin/sylphy

.PHONY: default
default: build run

.PHONY: build
build:
	@go build -o $(tmp_bin) -race $(entrypoint)

.PHONY: run
run:
	@$(tmp_bin)

.PHONY: test
test:
	@go test ./... -race