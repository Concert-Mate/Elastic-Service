package internal

import (
	"bytes"
	"context"
	"elastic-service/internal/helpers"
	pb "elastic-service/pkg/api"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"net/http"
)

func (ec *ElasticsearchClient) IndexDocument(indexName string, city *pb.City) error {

	// Serialize the city object to JSON
	body, err := json.Marshal(city)
	if err != nil {
		return err
	}

	// Create the request to index the document
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: city.Name,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	// Execute the request to index the document
	res, err := req.Do(context.Background(), ec.Client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	// Check for errors in the response
	if res.IsError() {
		var body map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return err
		}
		return fmt.Errorf("error indexing document: %s: %s", res.Status(), body["error"].(map[string]interface{})["reason"])
	}
	return nil
}

func (ec *ElasticsearchClient) initCities() error {
	check, err := checkNotExistsIndex(ec)
	if err != nil {
		return err
	}

	if !check {
		return nil
	}

	data, err := helpers.GetContentFromURL(fileNameUrl)
	if err != nil {
		return err
	}

	var cities []*pb.City

	err = json.Unmarshal(data, &cities)
	if err != nil {
		return err
	}

	// Define the mapping for the index
	mapping := `{
        "mappings": {
            "properties": {
                "name": { "type": "text" },
                "coords": { "type": "geo_point" }
            }
        }
    }`
	// Create the request to create the index with mapping
	createIndexReq := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  bytes.NewReader([]byte(mapping)),
	}

	// Execute the request to create the index
	createIndexRes, err := createIndexReq.Do(context.Background(), ec.Client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(createIndexRes.Body)

	if createIndexRes.IsError() {
		var body map[string]interface{}
		if err := json.NewDecoder(createIndexRes.Body).Decode(&body); err != nil {
			return err
		}
		return fmt.Errorf("error creating index: %s: %s", createIndexRes.Status(), body["error"].(map[string]interface{})["reason"])
	}

	// Index documents
	for _, city := range cities {
		err = ec.IndexDocument(indexName, city)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func checkNotExistsIndex(esClient *ElasticsearchClient) (bool, error) {
	resp, err := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}.Do(context.Background(), esClient.Client)

	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusNotFound, nil
}
