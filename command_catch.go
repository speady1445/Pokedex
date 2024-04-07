package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, inputs []string) error {
	if len(inputs) < 1 {
		return errors.New("missing pokemon name")
	}

	pokemon, err := cfg.pokeApi.GetPokemon(inputs[0])
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at", pokemon.Name, "...")
	roll := rand.Intn(int(math.Pow(float64(pokemon.BaseExperience), 1.2)))
	if roll > pokemon.BaseExperience {
		fmt.Println(pokemon.Name, "escaped!")
		return nil
	}

	fmt.Println(pokemon.Name, "caught!")
	cfg.caughtPokemons[pokemon.Name] = pokemon
	return nil
}
