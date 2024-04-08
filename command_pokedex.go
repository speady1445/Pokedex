package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ []string) error {
	if len(cfg.caughtPokemons) == 0 {
		fmt.Println("you have not caught any pokemons")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
