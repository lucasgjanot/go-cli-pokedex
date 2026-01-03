package commands

import (
	"errors"
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandExplore(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location_name>")
	}
	name := args[0]
	location, err := cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s", location.Name)
	Dotwait(3)
	println("\nFound Pokemon:")

	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}