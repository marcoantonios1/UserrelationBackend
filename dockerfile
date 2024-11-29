FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY cmd/app/.env .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin ./cmd/app/main.go

EXPOSE 9013

ENTRYPOINT [ "/app/bin" ]