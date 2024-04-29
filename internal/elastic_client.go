package internal

import (
	"context"
	pb "elastic-service/pkg/api"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"os"
	"strings"
)

type ElasticsearchClient struct {
	Client *elasticsearch.Client
}

func NewElasticsearchClient() (*ElasticsearchClient, error) {
	host := os.Getenv(elasticSearchAddress)
	if host == "" {
		host = defaultHost
	}

	cfg := elasticsearch.Config{
		Addresses: []string{defaultHost},
		Username:  os.Getenv(elasticUsername),
		Password:  os.Getenv(elasticPassword),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	res := &ElasticsearchClient{Client: es}

	err = res.initCities()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// SearchByName searches for a city by name
func (ec *ElasticsearchClient) SearchByName(name string) ([]*pb.City, error) {
	query := fmt.Sprintf(`{
		"query": {
			"match": {
				"name": {
					"query": "%s",
					"fuzziness": "%s"
				}
			}
		}
	}`, name, defaultFuzziness)

	return ec.search(query)
}

// SearchByCoords searches for a city by coordinates
func (ec *ElasticsearchClient) SearchByCoords(coords *pb.Coords) ([]*pb.City, error) {
	distance := os.Getenv(geoDistance)

	if distance == "" {
		distance = defaultDistance
	}

	query := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{
						"match_all": {}
					},
					{
						"geo_distance": {
							"distance": "%s",
							"coords": {
								"lat": %f,
								"lon": %f
							}
						}
					}
				]
			}
		}
	}`, distance, coords.GetLat(), coords.GetLon())

	return ec.search(query)
}

// search performs a generic search in Elasticsearch
func (ec *ElasticsearchClient) search(query string) ([]*pb.City, error) {
	res, err := ec.Client.Search(
		ec.Client.Search.WithContext(context.Background()),
		ec.Client.Search.WithIndex(indexName),
		ec.Client.Search.WithBody(strings.NewReader(query)),
		ec.Client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	var result struct {
		Hits struct {
			Hits []*struct {
				Source *pb.City `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	var cities []*pb.City
	for _, hit := range result.Hits.Hits {
		cities = append(cities, hit.Source)
	}

	return cities, nil
}
