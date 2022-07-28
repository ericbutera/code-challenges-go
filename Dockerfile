FROM golang:1.18-alpine
RUN apk add build-base

WORKDIR /app

COPY . ./

RUN go mod download

ENTRYPOINT [ "go", "test", "-v", "./..." ]
