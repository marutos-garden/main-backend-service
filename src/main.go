package main

import (
	"log"
	"net"

	pb "github.com/marutos/protobuf-files/device_service"
	"google.golang.org/grpc"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.DeviceServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Println("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterDeviceServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)

	}
}
