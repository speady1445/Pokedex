package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, inputs []string) error {
	if len(inputs) < 1 {
		return errors.New("missing area parameter")
	}

	area := inputs[0]
	pokemons, err := cfg.pokeApi.GetPokemonsInLocation(area)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", area, "...")
	for _, pokemon := range pokemons {
		fmt.Println(" - ", pokemon)
	}
	return nil
}
