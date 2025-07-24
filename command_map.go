package main

import "fmt"

func commandMap(cfg *config, areaName string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil

}

func commandMapB(cfg *config, areaName string) error {
	if cfg.prevLocationURL == nil {
		fmt.Println("You are on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
