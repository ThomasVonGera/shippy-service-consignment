package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ThomasVonGera/shippy-service-consignment/proto/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

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

type service struct {
	repo repositoryinterface
}

func (s *service) CreateConsignment(ctx context.Context, request *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(request)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, request *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fehler beim Listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterShippingServiceServer(server, &service{repo})

	reflection.Register(server)

	log.Println("Running on Port: ", port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Fehler beim Serven: %v", err)
	}
}
