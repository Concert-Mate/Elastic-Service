package handlers

import (
	"elastic-service/internal"
	pb "elastic-service/pkg/api"
	"log"
)

type Server struct {
	pb.UnimplementedCitySearchServer
	client *internal.ElasticsearchClient
}

func NewServer() *Server {
	ec, err := internal.NewElasticsearchClient()

	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		client: ec,
	}
}
