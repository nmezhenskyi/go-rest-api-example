build:
	go build -o ./bin/restapi ./cmd/restapi

dev:
	go run ./cmd/restapi

run:
	./bin/restapi

clean:
	rm ./bin/restapi

compile: build run

.DEFAULT_GOAL := compile
