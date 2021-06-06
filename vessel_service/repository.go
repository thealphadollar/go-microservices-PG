package main

import (
	"context"

	pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerId   string `json:"owner_id"`
}

func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		OwnerId:   vessel.OwnerId,
		Available: vessel.Available,
	}
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		OwnerId:   vessel.OwnerId,
		Available: vessel.Available,
	}
}

type repository interface {
	Create(ctx context.Context, vessel *Vessel) error
	FindAvailable(ctx context.Context, maxWeight int32, capacity int32) (*Vessel, error)
}

type MongoRepository struct {
	vessel_collection *mongo.Collection
}

func (repo *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repo.vessel_collection.InsertOne(ctx, vessel)
	return err
}

func (repo *MongoRepository) FindAvailable(ctx context.Context, maxWeight int32, capacity int32) (*Vessel, error) {
	var vessel *Vessel
	err := repo.vessel_collection.FindOne(ctx, bson.D{
		{Key: "capacity", Value: bson.D{{Key: "$gt", Value: capacity}}},
		{Key: "max_weight", Value: bson.D{{Key: "$gt", Value: maxWeight}}},
		{Key: "available", Value: true},
	}).Decode(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}
