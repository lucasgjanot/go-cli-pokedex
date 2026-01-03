package commands

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandCatch(cfg *config.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon_name>")
	}

	pokemonName := args[0]

	pokemonResp, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s", pokemonResp.Name)
	Dotwait(3)
	if caught := TryCatchPokemon(pokemonResp.BaseExperience); caught {
		cfg.Pokedex[pokemonResp.Name] = pokemonResp
		fmt.Printf("\n%s was caught!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("\n%s escaped!\n", pokemonResp.Name)

	return nil

}

func TryCatchPokemon(BaseExperience int) bool {
	chance := math.Pow(100.0/float64(BaseExperience), 1.3)

	if chance > 1 {
		chance = 1
	}

	return rand.Float64() <= chance
}
