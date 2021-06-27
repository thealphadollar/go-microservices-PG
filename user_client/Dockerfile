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
    -installsuffix cgo -o user-client ./user_client

# Run Service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/user-client .

CMD ["./user-client"]