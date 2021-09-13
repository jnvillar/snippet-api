
build:
	go build -o bin/snippet main.go

run: build
	./bin/snippet serve

