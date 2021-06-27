package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
)

const (
	defaultFilename = "consignment_client/consignment.json"
	token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiMWFiY2M0MGItNjYyZS00YzEyLTkyNzgtMWVmMzlhZDE3MmYwIiwibmFtZSI6IlNoaXZhbSIsImNvbXBhbnkiOiJNZXJjYXJpIiwiZW1haWwiOiJzLWt1bWFyQG1lcmNhcmkuY29tIiwicGFzc3dvcmQiOiIkMmEkMTAkRnYyWXdPajVkNDhOaXMuaDdDcG1tLlFWY2NiVGZxRm43dVJnN2pDQ3RCSWFrZ3JLbDRsUy4ifSwiaXNzIjoidXNlci5zZXJ2aWNlIn0.L5nOM2vk0iQDCMMRa3InrEnp3zGVYRnWPo7iTgTHxUo"
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

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Token": token,
	})

	response, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("failed to get response: %v", err)
	}
	log.Printf("Created: %t", response.Created)
	response2, err2 := client.CreateConsignment(ctx, consignment)
	if err2 != nil {
		log.Fatalf("failed to get response: %v", err2)
	}
	log.Printf("Created: %t", response2.Created)

	response3, err3 := client.GetConsignment(ctx, &pb.GetRequest{})
	if err3 != nil {
		log.Fatalf("failed to get consignments: %v", err3)
	}
	log.Println("Total consignments received: ", response3.TotalConsignments)
	for i, v := range response3.Consignments {
		log.Println(i, "consignment's description is", v.Description)
		log.Println(i, "consignment's vesselId is", v.VesselId)
	}
}
