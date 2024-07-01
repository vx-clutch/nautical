BIN_NAME=knot
MAIN=cmd/nautical/main.go

all: build

build:
	go build -o /bin/${BIN_NAME} ./...

test:
	go test -v ./...

clean:
	go clean
	rm ${BIN_NAME}
