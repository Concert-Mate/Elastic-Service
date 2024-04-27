package main

import (
	"elastic-service/internal/handlers"
	pb "elastic-service/pkg/api" // Import generated proto package
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Get port from command-line argument or use default
	port := ":50051"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCitySearchServer(s, handlers.NewServer())

	reflection.Register(s)

	log.Printf("gRPC server started on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
