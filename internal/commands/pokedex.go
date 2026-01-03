package commands

import (
	"errors"
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandPokedex(cfg *config.Config, args ...string) error {
	if len(args) != 0 {
		return errors.New("usage: pokedex")
	}

	if len(cfg.Pokedex) == 0 {
		return errors.New("Your Pokedex is empty!")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
