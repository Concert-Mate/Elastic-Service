package api

import "errors"

func NewCoords(lat, lon float32) (*Coords, error) {
	if lat < -90 || lat > 90 {
		return nil, errors.New("latitude out of range (-90 to 90 degrees)")
	}
	if lon < -180 || lon > 180 {
		return nil, errors.New("longitude out of range (-180 to 180 degrees)")
	}

	return &Coords{Lat: lat, Lon: lon}, nil
}
