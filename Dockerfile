# image for build
FROM golang:1.13.4-alpine3.10 as builder

ENV GO111MODULE on
ENV PROJECT_ROOT /go/src/github.com/istsh/go-grpc-sample/
WORKDIR $PROJECT_ROOT

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN apk add --no-cache mysql-client

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -tags gae -a -o build ./app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a ./app/cmd/client

# image for release
FROM gcr.io/distroless/base:latest as release
ENV BUILDER_ROOT /go/src/github.com/istsh/go-grpc-sample/
ENV PROJECT_ROOT /
COPY --from=builder $BUILDER_ROOT/build $PROJECT_ROOT/build
COPY --from=builder $BUILDER_ROOT/client $PROJECT_ROOT/client
