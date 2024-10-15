PROTO_DIR = internal/proto
OUT_DIR = internal/proto

PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto
proto:
	@echo "Generating gRPC code..."
	@protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) --proto_path=$(PROTO_DIR) $(PROTO_FILES)
	@echo "gRPC code generation complete."

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@rm -rf $(OUT_DIR)/*.go
	@echo "Cleanup complete."

.PHONY: run
run:
	@go run cmd/api/main.go

.PHONY: build
build:
	@go build -o bin/timeseries cmd/api/main.go
