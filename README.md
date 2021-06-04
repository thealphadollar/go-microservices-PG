# go-microservices-PG
A playground to learn making microservices in GO making use of protobuf and gRPC as the underlying transport protocol.

## Tech Stack

The tech stack includes golang, mongodb, grpc, docker, Google Cloud, Kubernetes, NATS, CircleCI, Terraform and go-micro.

## Dockerisation

We have used multistage builds to keep the size of the containers minimal.

### Service Docker

The dockerfile for service is present in the "service" folder in the project root.

0. `cd /root/of/the/project`
0. `docker build -t consignment-service -f ./service/Dockerfile .`
0. `docker run -p 50051:50051 consignment-service:latest`

### Client Docker

The dockerfile for client is present in the "client" folder in the project root.

0. `cd /root/of/the/project`
0. `docker build -t consignment-client -f ./client/Dockerfile .`
0. `docker run -p 50051:50051 consignment-client:latest`

## Credit

The playground follows the [10-part tutorial on building microservices in Go](https://ewanvalentine.io/microservices-in-golang-part-1/) by @EwanValentine.