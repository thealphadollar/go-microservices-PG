package main

import (
	"context"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
	vessel_pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	s := micro.NewService(
		micro.Name("consignment.service"),
	)
	s.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.TODO(), uri, 3)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consigments")
	repository := &MongoRepository{consignment_collection: consignmentCollection}
	vesselClient := vessel_pb.NewVesselService("vessel.service", s.Client())

	// add dummy vessels
	resp, err := vesselClient.CreateVessel(context.Background(), &vessel_pb.Vessel{
		Id: "001", Name: "Boaty Vessel 1", Capacity: 500, MaxWeight: 200000, Available: true,
	})
	if err != nil {
		log.Panicf("failed to create dummy vessel: %v", err)
	} else {
		log.Println("vessel creation: ", resp.Created)
	}

	h := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(s.Server(), h)
	if err := s.Run(); err != nil {
		log.Fatalf("failed to run service: %v", err)
	}
}
