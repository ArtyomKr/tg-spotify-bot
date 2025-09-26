FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/server ./cmd/


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

RUN mkdir -p /app/data

VOLUME /app/data

EXPOSE ${PORT}

CMD ["./server"]