FROM golang:1.21.0-alpine3.17 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o currency-exchange ./cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/currency-exchange .

# Command to run the executable
CMD ["./currency-exchange"]
