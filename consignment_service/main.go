package main

import (
	"context"
	"log"
	"sync"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/consignment_service/proto/consignment"
	vessel_pb "github.com/thealphadollar/go-microservices-PG/vessel_service/proto/vessel"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
}

// a dummy repository to simulate warehouse
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() ([]*pb.Consignment, error) {
	return repo.consignments, nil
}

type service struct {
	repo           repository
	vessel_service vessel_pb.VesselService
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vessel, err := s.vessel_service.FindAvailable(context.Background(), &vessel_pb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}

	req.VesselId = vessel.Vessel.Id

	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Consignment = consignment
	res.Created = true
	return nil
}

func (s *service) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	res.TotalConsignments = int32(len(consignments))
	return nil
}

func main() {
	repo := &Repository{}

	s := micro.NewService(
		micro.Name("consignment.service"),
	)
	s.Init()
	vessel_service := vessel_pb.NewVesselService("vessel.service", s.Client())

	// tie with grpc generated server
	if err := pb.RegisterShippingServiceHandler(s.Server(), &service{repo, vessel_service}); err != nil {
		log.Fatalf("failed to register service handler: %v", err)
	}

	log.Println("Running on port: ", port)
	if err := s.Run(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
