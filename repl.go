package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		userInput := cleanInput(text)
		commandName := userInput[0]
		locationName := ""
		if len(userInput) > 1 {
			locationName = userInput[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, locationName)
			if err != nil {
				fmt.Printf("Command '%s' failed: %s\n", commandName, err)
			}
		} else {
			fmt.Printf("Unknown Command\n")
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		}, "help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		}, "map": {
			name:        "map",
			description: "Get the next page",
			callback:    commandMap,
		}, "mapb": {
			name:        "mapb",
			description: "Get the previous page",
			callback:    commandMapB,
		}, "explore": {
			name:        "explore",
			description: "Explore what pokemon exist in the region",
			callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
