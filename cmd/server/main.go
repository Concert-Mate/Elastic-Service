package main

import (
	"elastic-service/internal"
	"elastic-service/internal/handlers"
	pb "elastic-service/pkg/api" // Import generated proto package
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Get port from command-line argument or use default
	port := internal.DefaultPort

	if len(os.Args) > 1 {
		portStr := os.Args[1]
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				log.Fatalf("Invalid port number: %s", portStr)
			}
		}
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCitySearchServer(s, handlers.NewServer())

	reflection.Register(s)

	log.Printf("gRPC server started on port %d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
