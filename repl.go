package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andybzn/pokedex/internal/pokeapi"
)

type config struct {
	ApiClient   pokeapi.Client
	NextUrl     *string
	PreviousUrl *string
}

var cfg = &config{
	ApiClient: pokeapi.NewClient(10*time.Second, 5*time.Minute),
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		if len(command) == 0 {
			continue
		}
		c, exists := cliCommands()[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := c.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List the next 20 locations on the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 locations on the map",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore the specified location area, listing the available Pokemon",
			callback:    commandExplore,
		},
	}
}
