package main

import (
	"context"
	"log"

	pb "github.com/marutos/protobuf-files/device_service"
)

func (s *Server) Device(ctx context.Context, in *pb.DeviceRequest) (*pb.DeviceResponse, error) {
	log.Printf("Device function was invoked with %v\n", in)
	return &pb.DeviceResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
