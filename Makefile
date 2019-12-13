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

generate:
	GO111MODULE=on \
	go generate ./...

run:
	docker-compose -f docker-compose.local.yml up --build client server

stop:
	docker-compose -f docker-compose.local.yml stop client server

db-run:
	docker-compose -f docker-compose.local.yml up -d db

db-stop:
	docker-compose -f docker-compose.local.yml stop db
