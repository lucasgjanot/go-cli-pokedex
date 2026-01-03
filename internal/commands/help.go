package commands

import (
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandHelp(cfg *config.Config, args ...string) error {
	commandsRegistry := GetCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range commandsRegistry {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
