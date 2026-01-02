package main

import (
	"time"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
	"github.com/lucasgjanot/go-cli-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config.Config{
		PokeapiClient: &pokeClient,
	}
	startRepl(cfg)
}
