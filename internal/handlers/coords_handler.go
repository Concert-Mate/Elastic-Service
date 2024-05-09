package handlers

import (
	"context"
	pb "elastic-service/pkg/api"
)

func (s *Server) SearchByCoords(ctx context.Context, req *pb.CoordsRequest) (*pb.CitySearchCoordsResponse, error) {
	response := &pb.CitySearchCoordsResponse{}

	lat := req.GetLat()
	lon := req.GetLon()
	radius := req.GetRadius()

	coords, err := pb.NewCoords(lat, lon)
	if err != nil {
		response.Code = pb.CitySearchCoordsResponse_INVALID_COORDS
		return response, nil
	}

	cities, err := s.client.SearchByCoords(coords, radius)
	if err != nil {
		response.Code = pb.CitySearchCoordsResponse_INTERNAL_ERROR_COORDS
		return response, nil
	}

	response.Options = cities
	if len(cities) == 0 {
		response.Code = pb.CitySearchCoordsResponse_EMPTY_COORDS
	} else {
		response.Code = pb.CitySearchCoordsResponse_SUCCESS_COORDS
	}
	return response, nil
}
