gen-proto:
	@echo "  >  Generate Protobuf..."
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

gen-wire:
	@echo "  >  Generate Wire..."
	# wire wire/wire.go

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