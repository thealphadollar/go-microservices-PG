package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if (spec.Capacity <= vessel.Capacity) && (spec.MaxWeight <= vessel.MaxWeight) && vessel.Available {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found matching given specifications!")
}

type VesselService struct {
	repo Repository
}

func (s *VesselService) FindAvailable(ctx context.Context, specs *pb.Specification, res *pb.Response) error {
	if vessel, err := s.repo.FindAvailable(specs); err != nil {
		return err
	} else {
		res.Vessel = vessel
		return nil
	}
}

func main() {
	vessels := []*pb.Vessel{
		{Id: "001", Name: "Boaty Vessel 1", Capacity: 500, MaxWeight: 200000, Available: true},
		{Id: "002", Name: "Boaty Vessel 2", Capacity: 100, MaxWeight: 200000, Available: false},
	}
	repo := &VesselRepository{vessels: vessels}
	vesselService := VesselService{repo}

	service := micro.NewService(
		micro.Name("vessel.service"),
	)
	service.Init()

	if err := pb.RegisterVesselServiceHandler(service.Server(), &vesselService); err != nil {
		log.Fatalf("failed to register vessel service!")
	}

	if err := service.Run(); err != nil {
		log.Fatalf("failed to run service!")
	}
}
