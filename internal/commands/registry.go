package commands

import (
	"fmt"
	"time"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch a pokemon",
			Callback:    CommandCatch,
		},
	}
}

func Dotwait(nDots int) {
	for i := 0; i < nDots; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
}
