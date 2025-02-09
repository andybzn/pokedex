package main

import "fmt"

func commandMap(cfg *config) error {
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

func commandMapB(cfg *config) error {
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
