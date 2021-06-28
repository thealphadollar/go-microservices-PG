package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	k8s "github.com/micro/kubernetes/go/micro"
	pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	service := k8s.NewService(
		micro.Name("vessel.service"),
	)
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(context.TODO(), uri, 3)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vessel_collection: vesselCollection}

	h := &handler{
		repo: repository,
	}

	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Fatalf("failed to register vessel service!")
	}

	if err := service.Run(); err != nil {
		log.Fatalf("failed to run service!")
	}
}
