package commands

import (
	"fmt"
	"os"

	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandExit(cfg *config.Config, args ...string) error {
	fmt.Print("Closing the Pokedex")
	Dotwait(3)
	fmt.Println(" Goodbye!")
	os.Exit(0)
	return nil
}
