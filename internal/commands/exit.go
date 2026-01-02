package commands

import (
	"fmt"
	"os"
	"time"
	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func CommandExit(cfg *config.Config) error {
	fmt.Print("Closing the Pokedex")
	dotwait(3)
	fmt.Println(" Goodbye!")
	os.Exit(0)
	return nil
}

func dotwait(nDots int) {
	for i := 0; i < nDots; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
}