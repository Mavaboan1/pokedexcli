package main

import "fmt"

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	height := pokemon.Height
	weight := pokemon.Weight
	stats := pokemon.Stats
	hp := stats[0].BaseStat
	attack := stats[1].BaseStat
	defense := stats[2].BaseStat
	specialAttack := stats[3].BaseStat
	specialDefense := stats[4].BaseStat
	speed := stats[5].BaseStat
	types := pokemon.Types
	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Weight: %d\n", weight)
	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", hp)
	fmt.Printf("  -attack: %d\n", attack)
	fmt.Printf("  -defense: %d\n", defense)
	fmt.Printf("  -special-attack: %d\n", specialAttack)
	fmt.Printf("  -special-defense: %d\n", specialDefense)
	fmt.Printf("  -speed: %d\n", speed)
	fmt.Println("Types:")
	for _, t := range types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	return nil
}
