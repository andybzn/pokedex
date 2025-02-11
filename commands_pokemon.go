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
		cfg.PokeDex.Pokemon[pokemon] = data
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
