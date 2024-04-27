package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (c *Coords) UnmarshalJSON(data []byte) error {
	// Структура для временного хранения значений координат как interface{}
	type coordsJSON struct {
		Lat interface{} `json:"lat"`
		Lon interface{} `json:"lon"`
	}
	var temp coordsJSON

	// Десериализуем данные во временную структуру
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Функция для преобразования interface{} в float32
	convertToFloat32 := func(val interface{}) (float32, error) {
		switch v := val.(type) {
		case float64:
			return float32(v), nil
		case string:
			f, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return 0, err
			}
			return float32(f), nil
		default:
			return 0, fmt.Errorf("unsupported type for coordinates")
		}
	}

	// Преобразуем и присваиваем значения координат
	var err error
	if c.Lat, err = convertToFloat32(temp.Lat); err != nil {
		return err
	}
	if c.Lon, err = convertToFloat32(temp.Lon); err != nil {
		return err
	}

	return nil
}
