package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/thealphadollar/go-microservices-PG/service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "client/consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("failed to read file: %v", err)
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client := pb.NewShippingServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("failed to parse file: %v", err)
	}

	response, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("failed to get response: %v", err)
	}
	log.Printf("Created: %t", response.Created)
	log.Printf("Consignment Description: %s", response.Consignment.Description)
}
