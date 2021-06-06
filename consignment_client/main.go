package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
)

const (
	defaultFilename = "consignment_client/consignment.json"
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
	s := micro.NewService(
		micro.Name("consignment.client"),
	)
	client := pb.NewShippingService("consignment.service", s.Client())

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
		log.Println(i, "consignment's vesselId is", v.VesselId)
	}
}
