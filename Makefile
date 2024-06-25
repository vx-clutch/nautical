BIN_NAME=main.out

build:
	go build -o ${BIN_NAME} ./...

run:
	go build -o ${BIN_NAME} ./...
	./${BIN_NAME}

install:
	go build -o /usr/bin/knot ./...
