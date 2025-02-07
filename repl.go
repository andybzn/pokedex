package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		command := cleanInput(scanner.Text())[0]
		if len(command) == 0 {
			continue
		}
		c, exists := commandRegistry[command]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := c.callback()
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
	callback    func() error
}

var commandRegistry = map[string]cliCommand{}

func init() {
	commandRegistry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the program",
		callback:    commandExit,
	}
	commandRegistry["help"] = cliCommand{
		name:        "help",
		description: "Display this help message",
		callback:    commandHelp,
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range commandRegistry {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
