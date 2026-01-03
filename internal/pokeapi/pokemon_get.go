package pokeapi

import "encoding/json"

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	var zero Pokemon
	var pokemonResp Pokemon
	
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return zero, err
		}

		return pokemonResp, nil
	}

	
	data, err := GetData(url)
	if err != nil {
		return zero, err
	}

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return zero, err
	}

	c.cache.Add(url, data)
	
	return  pokemonResp, nil
}