.PHONY: build run

build:
	@echo "Building currency exchange service..."
	@docker build -t currency-exchange:latest .

run: build
	@echo "Running currency exchange service..."
	@docker-compose up -d
