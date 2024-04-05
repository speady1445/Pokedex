package main

import (
	"fmt"
)

func commandMapForward(cfg *config) error {
	return traverse(cfg.pokeApi.Map, "You are on the last page")
}

func commandMapBack(cfg *config) error {
	return traverse(cfg.pokeApi.Mapb, "You are on the first page")
}

func traverse(traverseFunction func() ([]string, error), noLocationsMessage string) error {
	locations, err := traverseFunction()
	if err != nil {
		return err
	}

	if len(locations) == 0 {
		fmt.Println(noLocationsMessage)
		return nil
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
