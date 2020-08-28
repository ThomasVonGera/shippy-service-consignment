package main

import (
	"context"
	"log"

	pb "github.com/ThomasVonGera/shippy-service-consignment/proto/consignment"
	"github.com/micro/go-micro/v2"
)

const serviceName = "shippy-.service.Consignment"

type repositoryinterface interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}
type Repository struct {
	consignments []*pb.Consignment
}

// Neue Consignment erstellen
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo repositoryinterface
}

func (s *consignmentService) CreateConsignment(ctx context.Context, request *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(request)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, request *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	service := micro.NewService(
		micro.Name(serviceName),
	)

	// initialisieren
	service.Init()

	//registrieren
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	//server ausführen
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
