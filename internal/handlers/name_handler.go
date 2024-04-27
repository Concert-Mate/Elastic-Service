package handlers

import (
	"context"
	pb "elastic-service/pkg/api"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Server) SearchByName(ctx context.Context, req *pb.CityNameRequest) (*pb.CitySearchNameResponse, error) {
	response := &pb.CitySearchNameResponse{}

	name := req.GetName()

	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is empty")
	}

	cities, err := s.client.SearchByName(name)
	if err != nil {
		fmt.Println(err)
		response.Code = pb.CitySearchNameResponse_INTERNAL_ERROR_NAME
		return response, nil
	}

	if len(cities) == 0 {
		// No cities found
		return &pb.CitySearchNameResponse{
			Code: pb.CitySearchNameResponse_EMPTY_NAME,
		}, nil
	}

	foundCityName := cities[0].Name

	code := pb.CitySearchNameResponse_FUZZY_NAME

	if strings.ToLower(foundCityName) == strings.ToLower(name) {
		code = pb.CitySearchNameResponse_SUCCESS_NAME
	}

	// Populate response with cities and status code
	response.Options = cities
	response.Code = code

	return response, nil
}
