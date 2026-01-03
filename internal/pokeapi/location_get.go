package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	var zero Location
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return zero, err
		}

		return locationResp, nil
	}

	data, err := GetData(url)
	if err != nil {
		return zero, err
	}

	var deepLocations Location
	err = json.Unmarshal(data, &deepLocations)
	if err != nil {
		return zero, err
	}

	c.cache.Add(url, data)

	return deepLocations, nil
}
