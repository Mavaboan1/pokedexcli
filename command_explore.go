package main

import "fmt"

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		return fmt.Errorf("area cannot be empty")
	}
	areaResp, err := cfg.pokeapiClient.ListAreaEncounter(areaName)
	if err != nil {
		return fmt.Errorf("Area '%s' does not exist", areaName)
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon")
	for _, pokemon := range areaResp.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}
	return nil
}
