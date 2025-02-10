package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, v := range cliCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}
