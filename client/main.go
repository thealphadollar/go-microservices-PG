package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	pb "github.com/thealphadollar/go-microservices-PG/service/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
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
	conn, err := grpc.Dial(address, grpc.WithInsecure)
	defer conn.Close()
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client = pb.NewShippingServiceClient(conn)
	
}
