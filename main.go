package main

import (
	"context"
	"log"

	pb "github.com/ThomasVonGera/shippy-service-consignment/proto/consignment"
	vesselProto "gitbug.com/ThomasVonGera/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
		serviceName = "shippy-.service.Consignment"
		defaultHost = "datastore:27017"
)


func main() {
	service := micro.NewService(
		micro.Name(serviceName)
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client,err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}

	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	vesselClient := vesselProto.NewVesselService("shippy.service.client", service.Client())

	hilfsHandler := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(service.Server(), hilfsHandler)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
