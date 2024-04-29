package internal

import (
	pb "elastic-service/pkg/api"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

var ecClient *ElasticsearchClient

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func setup() error {
	rootDir := RootDir()

	envFilePath := filepath.Join(rootDir, ".env")
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error loading .env file:", err)
		return err
	}

	var err error
	ecClient, err = NewElasticsearchClient()
	if err != nil {
		return err
	}

	return nil
}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		panic(err)
	}
	defer os.Clearenv()
	m.Run()
}

func TestSearchMoscowExists(t *testing.T) {
	cities, err := ecClient.SearchByName("Москва")
	assert.NoError(t, err, "Search by name failed")
	assert.NotEmpty(t, cities, "Expected at least one city, got none")
}

func TestSearchMoscowFirstFoundCorrect(t *testing.T) {
	cities, _ := ecClient.SearchByName("Москва")
	assert.Equal(t, "Москва", cities[0].Name)
}

func TestSearchMoscowFuzzy(t *testing.T) {
	cities, err := ecClient.SearchByName("Масква")
	assert.NoError(t, err, "Search by name failed")
	assert.NotEmpty(t, cities, "Expected at least one city, got none")

	foundExact := false
	for _, city := range cities {
		if city.Name == "Москва" {
			foundExact = true
			break
		}
	}
	assert.True(t, foundExact, "Fuzzy match for 'Москва' not found")
}

func TestSearchMoscowByCoords(t *testing.T) {
	moscowCoords := &pb.Coords{
		Lat: 55.755833333333,
		Lon: 37.617777777778,
	}

	cities, err := ecClient.SearchByCoords(moscowCoords)
	assert.NoError(t, err, "Search by coords failed")
	assert.NotEmpty(t, cities, "Expected at least one city, got none")

	// Check if Moscow is within the search radius
	moscowFound := false
	for _, city := range cities {
		if city.Name == "Москва" {
			moscowFound = true
			break
		}
	}
	assert.True(t, moscowFound, "Moscow not found within the specified radius")
}

func TestSearchNonExistentCityByName(t *testing.T) {
	cities, err := ecClient.SearchByName("НесуществующийГород")
	assert.NoError(t, err, "Search by non-existent city name failed")
	assert.Empty(t, cities, "Expected empty result for non-existent city")
}

func TestSearchByNonExistentCoords(t *testing.T) {
	nonExistentCoords := &pb.Coords{
		Lat: 1000.0,
		Lon: 1000.0,
	}

	cities, err := ecClient.SearchByCoords(nonExistentCoords)
	assert.NoError(t, err, "Search by non-existent coordinates failed")
	assert.Empty(t, cities, "Expected empty result for non-existent coordinates")
}

func TestConstructorValidation(t *testing.T) {
	_, err := pb.NewCoords(1000, 2000)
	assert.Error(t, err, "Expected error for invalid coords when create coords object by constructor")
}

func TestConstructorValidationNegative(t *testing.T) {
	_, err := pb.NewCoords(-1, -1)
	assert.Equal(t, err, nil)
}
