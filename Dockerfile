FROM golang:1.20-alpine AS builder

RUN apk update && \
    apk add --no-cache git make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest

# RUN apk update && apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]