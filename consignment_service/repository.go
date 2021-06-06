package main

import (
	"context"
	"log"

	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type Consignment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
	Origin string `json:"origin"`
}

type Containers []*Container

func MarshalContainerCollection(containers []*pb.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

func UnmarshalContainerCollection(containers []*Container) []*pb.Container {
	collection := make([]*pb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

func MarshalConsignmentCollection(consignments []*pb.Consignment) []*Consignment {
	collection := make([]*Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, MarshalConsignment(consignment))
	}
	return collection
}

func UnmarshalConsignmentCollection(consignments []*Consignment) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}
	return collection
}

func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID: container.Id,
		CustomerID: container.CustomerId,
		UserID: container.UserId,
		Origin: container.Origin,
	}
}

func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id: container.ID,
		CustomerId: container.CustomerID,
		UserId: container.UserID,
		Origin: container.Origin,
	}
}

func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	return &Consignment{
		ID: consignment.Id,
		Weight: consignment.Weight,
		VesselID: consignment.VesselId,
		Containers: MarshalContainerCollection(consignment.Containers),
		Description: consignment.Description,
	}
}

func UnmarshalConsignment (consignment *Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id: consignment.ID,
		Description: consignment.Description,
		VesselId: consignment.VesselID,
		Weight: consignment.Weight,
		Containers: UnmarshalContainerCollection(consignment.Containers),
	}
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

type MongoRepository struct {
	consignment_collection *mongo.Collection
}

func (repo *MongoRepository) Create (ctx context.Context, consignment *Consignment) error {
	_, err := repo.consignment_collection.InsertOne(ctx, consignment)
	return err
}

func (repo *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := repo.consignment_collection.Find(ctx, nil, nil)
	if err != nil {
		log.Panicf("failed to find in repo collection: %v", err)
	}
	var consignments []*Consignment
	for cur.Next(ctx) {
		var consignment *Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, nil
}