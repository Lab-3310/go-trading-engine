build:
	go build -o bin/orderbook

run: build
	./bin/orderbook

test:
	go test -v ./...
