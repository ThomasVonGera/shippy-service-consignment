package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/ThomasVonGera/shippy-service-consignment/proto/consignment"

	//"github.com/grpc-go/reflection"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repositoryinterface interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Neue Consignment erstellen
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

type service struct {
	repo repositoryinterface
}

func (s *service) CreateConsignment(ctx context.Context, request *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repo.Create(request)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: consignment}, nil
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
