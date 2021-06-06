package main

import (
	"context"

	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
	vessel_pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

type handler struct {
	repo         repository
	vesselClient vessel_pb.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vessel_pb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := s.repo.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *handler) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	res.TotalConsignments = int32(len(res.Consignments))
	return nil
}
