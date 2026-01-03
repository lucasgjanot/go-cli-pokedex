package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lucasgjanot/go-cli-pokedex/internal/commands"
	"github.com/lucasgjanot/go-cli-pokedex/internal/config"
)

func StartRepl(cfg *config.Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
