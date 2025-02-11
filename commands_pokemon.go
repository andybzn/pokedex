package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("nothing to catch!")
	}
	pokemon := args[0]
	data, err := cfg.ApiClient.FetchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	catchChance := rand.Intn(data.BaseExperience)
	if catchChance < 45 {
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Println("You may now inspect it with the `inspect` command")
		cfg.PokeDex.Pokemon[pokemon] = data
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("nothing to inspect!")
	}

	pokemon := args[0]
	pokemonData, exists := cfg.PokeDex.Pokemon[pokemon]
	if !exists {
		return fmt.Errorf("%s has not been caught yet!", pokemon)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemonData.Name, pokemonData.Height, pokemonData.Weight)
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemonData.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}
	fmt.Println("Abilities:")
	for _, abilities := range pokemonData.Abilities {
		fmt.Printf("  - %s\n", abilities.Ability.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.PokeDex.Pokemon

	if len(pokedex) == 0 {
		return fmt.Errorf("you have not caught any Pokemon yet!")
	}

	fmt.Println("Your Pokedex:")
	for key := range pokedex {
		fmt.Printf("  - %s\n", key)
	}

	return nil
}
