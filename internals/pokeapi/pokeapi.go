package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeAPI struct {
	next     *string
	previous *string
}

func GetPokeAPI() *PokeAPI {
	startPlace := "https://pokeapi.co/api/v2/location-area/"
	return &PokeAPI{
		next:     &startPlace,
		previous: nil,
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

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *PokeAPI) getLocations(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}

	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if response.StatusCode > 299 {
		return []string{}, fmt.Errorf("response failed with status code: %d", response.StatusCode)
	}
	if err != nil {
		return []string{}, err
	}

	decodedResponse := locationAreaResponse{}
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
