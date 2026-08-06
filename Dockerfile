FROM golang:1.26.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o renewit-api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/renewit-api .

EXPOSE 8080

CMD ["./renewit-api"]