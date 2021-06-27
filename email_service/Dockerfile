FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY ./go.mod ./go.mod

RUN go mod download

COPY ./ ./


RUN CGO_ENABLED=0 GOOS=linux go build -a \
    -installsuffix cgo -o email-service ./email_service

# Run Service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/email-service .

CMD ["./email-service"]