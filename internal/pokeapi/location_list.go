package pokeapi

import (
	"encoding/json"
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
			return zero, err
		}

		return locationsResp, nil
	}

	data, err := GetData(url)
	if err != nil {
		return zero, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return zero, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
