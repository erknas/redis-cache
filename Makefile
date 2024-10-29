build:
	@go build -o bin/redis-cache

run: build
	@./bin/redis-cache