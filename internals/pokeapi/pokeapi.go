package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/speady1445/Pokedex/internals/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type PokeAPI struct {
	next     *string
	previous *string
	cache    pokecache.Cache
}

func GetPokeAPI() *PokeAPI {
	startPlace := baseURL + "/location-area/"
	return &PokeAPI{
		next:     &startPlace,
		previous: nil,
		cache:    pokecache.NewCache(5 * time.Minute),
	}
}

func (c *PokeAPI) Map() ([]string, error) {
	if c.next == nil {
		return []string{}, nil
	}
	return c.getLocations(*c.next)
}

func (c *PokeAPI) Mapb() ([]string, error) {
	if c.previous == nil {
		return []string{}, nil
	}
	return c.getLocations(*c.previous)
}

type locationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *PokeAPI) getLocations(url string) ([]string, error) {
	body, err := c.fetchBody(url)
	if err != nil {
		return []string{}, err
	}

	decodedResponse := locationAreas{}
	err = json.Unmarshal(body, &decodedResponse)
	if err != nil {
		return []string{}, err
	}

	c.next = decodedResponse.Next
	c.previous = decodedResponse.Previous

	locations := make([]string, 0, len(decodedResponse.Results))
	for _, location := range decodedResponse.Results {
		locations = append(locations, location.Name)
	}
	return locations, nil
}

type locationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *PokeAPI) GetPokemonsInLocation(locationName string) ([]string, error) {
	body, err := c.fetchBody(baseURL + "/location-area/" + locationName)
	if err != nil {
		return []string{}, err
	}

	locationResponse := locationArea{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return []string{}, err
	}

	pokemons := make([]string, 0, len(locationResponse.PokemonEncounters))
	for _, pokemon := range locationResponse.PokemonEncounters {
		pokemons = append(pokemons, pokemon.Pokemon.Name)
	}
	return pokemons, nil
}

func (c *PokeAPI) fetchBody(url string) ([]byte, error) {
	body, ok := c.cache.Get(url)
	if ok {
		return body, nil
	}

	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err = io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status code: %d", response.StatusCode)
	}
	if err != nil {
		return []byte{}, err
	}

	c.cache.Add(url, body)
	return body, nil
}
