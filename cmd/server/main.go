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
	defaultPort := strconv.Itoa(internal.DefaultPort)

	configuredPort := os.Getenv("PORT")
	port := configuredPort
	if configuredPort == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", ":"+port)
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
