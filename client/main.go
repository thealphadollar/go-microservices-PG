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
	response2, err2 := client.CreateConsignment(context.Background(), consignment)
	if err2 != nil {
		log.Fatalf("failed to get response: %v", err2)
	}
	log.Printf("Created: %t", response2.Created)

	response3, err3 := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err3 != nil {
		log.Fatalf("failed to get consignments: %v", err)
	}
	log.Println("Total consignments received: ", response3.TotalConsignments)
	for i, v := range response3.Consignments {
		log.Println(i, "consignment's description is", v.Description)
	}
}
