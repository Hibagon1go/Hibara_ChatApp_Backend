FROM golang:1.18-alpine

WORKDIR /app/go/base

COPY go.mod .
COPY go.sum .

EXPOSE 8080
ENV GO111MODULE=on

RUN go mod download
RUN go install github.com/cosmtrek/air@latest
COPY . .

CMD ["go","run","cmd/app/main.go"]