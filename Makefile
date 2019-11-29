init:
	GO111MODULE=on \
	go mod download

build:
	GO111MODULE=on \
	go build

test:
	GO111MODULE=on \
	go test ./...

test-v:
	GO111MODULE=on \
	go test -v ./...

benchmark:
	GO111MODULE=on \
	go test -bench . -benchmem

lint:
	GO111MODULE=on \
	golint ./...

run:
	docker-compose -f ./docker-compose.yml up --build

stop:
	docker-compose -f ./docker-compose.yml stop

generate:
	GO111MODULE=on \
	go generate ./...
