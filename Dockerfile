FROM golang:1.15-alpine

RUN apk add --update git curl g++

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -v cmd/mock/mock.go

CMD [ "go", "test", "-v" ]
