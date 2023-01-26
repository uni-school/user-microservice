gen-proto:
	@echo "  >  Generate Protobuf..."
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user.proto


gen-wire:
	@echo "  >  Generate Wire..."
	wire wire/wire.go

# START APP
start-dev: gen-proto gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env dev

start-test: gen-proto gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env test

start-stag: gen-proto gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env stag

start-prod: gen-proto gen-wire
	@echo "  >  Starting Program..."
	go run cmd/main.go -env prod