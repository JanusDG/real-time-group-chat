check:
	go vet ./...

format:
	go fmt ./...
	
build:
	go build -o bin/main main.go

prerun: check format build

run:
	go run main.go

compile:
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

launch: prerun run

all: prerun open run
	