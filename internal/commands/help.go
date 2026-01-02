package commands

import (
	"fmt"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandHelp(cfg *config.Config) error {
	commandsRegistry := GetCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, value := range commandsRegistry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
