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
    -installsuffix cgo -o consignment-client ./client

# Run Service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
RUN mkdir /app/client/
COPY ./client/consignment.json ./app/client/
WORKDIR /app
COPY --from=builder /app/consignment-client .

CMD ["./consignment-client"]