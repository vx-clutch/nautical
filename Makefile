BIN_NAME=knot

all: build test

build:
	go build -o ${BIN_NAME} ./...

test:
	go test -v ./...

run:
	go build -o ${BIN_NAME} ./...
	./${BIN_NAME}

install:
	go build -o ./usr/bin/knot ./...

clean:
	go clean
	rm ${BIN_NAME}

compile:
	go build -o ./bin/knot ./...
