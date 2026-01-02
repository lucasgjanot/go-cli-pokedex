package config

import "github.com/lucasgjanot/go-cli-pokedex/internal/pokeapi"

type Config struct {
	PokeapiClient    *pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}