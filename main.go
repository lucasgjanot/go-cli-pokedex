package main

import (
	"time"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
	"github.com/lucasgjanot/go-cli-pokedex/internal/pokeapi"
	"github.com/lucasgjanot/go-cli-pokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config.Config{
		PokeapiClient: &pokeClient,
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}
	repl.StartRepl(cfg)
}
