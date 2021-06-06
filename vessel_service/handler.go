package main

import (
	"context"

	pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

type handler struct {
	repo repository
}

func (s *handler) FindAvailable(ctx context.Context, specs *pb.Specification, res *pb.Response) error {
	if vessel, err := s.repo.FindAvailable(ctx, specs.MaxWeight, specs.Capacity); err != nil {
		return err
	} else {
		res.Vessel = UnmarshalVessel(vessel)
		return nil
	}
}

func (s *handler) CreateVessel(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repo.Create(ctx, MarshalVessel(req)); err != nil {
		res.Created = false
		return err
	} else {
		res.Created = true
		res.Vessel = req
		return nil
	}
}
