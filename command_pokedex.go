package main

import "fmt"

func commandPokedex(cfg *config, pokemonName string) error {
	pokedex := cfg.caughtPokemon
	for _, pokemon := range pokedex {
		fmt.Println(" - " + pokemon.Name)
	}
	return nil
}
