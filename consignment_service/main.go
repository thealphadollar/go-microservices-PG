package main

import (
	"context"
	"errors"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
	user_pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
	vessel_pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	s := micro.NewService(
		micro.Name("consignment.service"),
		micro.WrapHandler(AuthWrapper),
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

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["Token"]

		authClient := user_pb.NewUserService("user.service", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &user_pb.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err2 := fn(ctx, req, resp)
		return err2
	}
}
