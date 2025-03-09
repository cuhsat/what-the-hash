.PHONY: all clean

all: build

test:
	go test -v -race ./...

build:
	mkdir -p ./bin
	go build -v -race -o ./bin/wth cmd/wth/main.go

clean:
	rm -rf ./bin
	go clean
