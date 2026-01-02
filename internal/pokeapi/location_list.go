package pokeapi

import (
	"encoding/json"
	"fmt"
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

	res, err := http.Get(url)
	if err != nil {
		return zero, fmt.Errorf("Error at making request: %w", err)
	}
	defer res.Body.Close()

	var locationsResp RespShallowLocations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationsResp); err != nil {
		return zero, fmt.Errorf("Error at decoding json: %w", err)
	}
	
	return locationsResp, nil
}
