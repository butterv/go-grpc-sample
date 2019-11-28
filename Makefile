init:
	GO111MODULE=on go mod download

build:
	GO111MODULE=on go build

test:
	go test ./...

test-v:
	go test -v ./...

benchmark:
	go test -bench . -benchmem

lint:
	GO111MODULE=on golint ./...