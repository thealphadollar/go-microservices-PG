# Build Service
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
    -installsuffix cgo -o consignment-client ./consignment_client

# Run Service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /app/consignment_client/
COPY ./consignment_client/consignment.json ./app/consignment_client/
WORKDIR /app
COPY --from=builder /app/consignment-client .

CMD ["./consignment-client"]