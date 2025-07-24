package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return fmt.Errorf("pokemon cannot be empty")
	}
	pokemonResp, err := cfg.pokeapiClient.ListPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	experience := pokemonResp.BaseExperience
	res := rand.Intn(experience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if res > 40 {
		fmt.Printf("%s escaped\n", pokemonName)
		return nil
	}

	fmt.Printf("%s was caught\n", pokemonName)

	cfg.caughtPokemon[pokemonName] = pokemonResp

	return nil
}
