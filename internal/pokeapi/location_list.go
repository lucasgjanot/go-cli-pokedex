package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	var zero RespShallowLocations

	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return zero, fmt.Errorf("Error at decoding cache json: %w", err)
		}

		return locationsResp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return zero, fmt.Errorf("Error at making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return zero, fmt.Errorf("Error at reading response body: %w", err)
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return zero, fmt.Errorf("Error at parsing response body: %w", err)
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
