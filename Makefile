install:
	@echo "Installing dependencies..."
	@go mod tidy

build:
	go build -o bin/radiant-case-api ./cmd/radiant-case-api

run: build
	./bin/radiant-case-api

doc:
	@echo "Generating OpenAPI v3.1 documentation..."
	@go run scripts/openapi/generate.go

clean:
	rm -rf bin docs/*
