# Build Service
FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a \
    -installsuffix cgo -o consignment-service ./service

# Run Service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

ENV MICRO_SERVER_ADDRESS=:50051

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/consignment-service .

CMD ["./consignment-service"]