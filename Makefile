generate_wire:
	@echo "~~ Generating wire from ~~~"
	@wire ./internal/injection

start: 
	@echo "@@@@@@~~~~~~ Starting Inventories Service ~~~~~~@@@@@@"
	@go run cmd/main.go