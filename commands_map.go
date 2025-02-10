package main

import "fmt"

func commandMap(cfg *config, args ...string) error {
	if cfg.NextUrl == nil && cfg.PreviousUrl != nil {
		return fmt.Errorf("you are already at the end of the location list")
	}

	data, err := cfg.ApiClient.FetchLocations(cfg.NextUrl)
	if err != nil {
		return err
	}
	cfg.NextUrl = data.Next
	cfg.PreviousUrl = data.Previous
	for _, location := range data.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.PreviousUrl == nil {
		return fmt.Errorf("you are already at the start of the location list")
	}

	data, err := cfg.ApiClient.FetchLocations(cfg.PreviousUrl)
	if err != nil {
		return err
	}
	cfg.NextUrl = data.Next
	cfg.PreviousUrl = data.Previous
	for _, location := range data.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you haven't provided a location to explore!")
	}
	area := args[0]
	fmt.Printf("Exploring %s...\n", area)
	data, err := cfg.ApiClient.ExploreLocation(area)
	if err != nil {
		return err
	}

	if len(data.Encounters) == 0 {
		return fmt.Errorf("no pokemon found!")
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range data.Encounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}
