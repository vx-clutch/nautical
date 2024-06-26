BIN_NAME=knot
MAIN=cmd/nautical/main.go

all: build

build:
	go build -o ${BIN_NAME} ${MAIN}

test:
	go test -v ./...

run:
	go build -o ${BIN_NAME} ${MAIN}
	./${BIN_NAME}

install:
	go build -o ./usr/bin/${BIN_NAME} ${MAIN}

clean:
	go clean
	rm ${BIN_NAME}

compile:
	go build -o ./bin/knot ${MAIN}
