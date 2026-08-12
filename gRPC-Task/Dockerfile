FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server-app ./server

FROM alpine

WORKDIR /app

COPY --from=builder /app/server-app .

CMD ["./server-app"]