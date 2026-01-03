package commands

import (
	"errors"
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandInspect(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon_name>")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return errors.New("you have not caught that Pokémon")
	}
	fmt.Print("Analysing Pokémon")
	Dotwait(3)
	fmt.Println("\nName:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}
