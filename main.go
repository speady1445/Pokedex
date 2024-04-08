package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/speady1445/Pokedex/internals/pokeapi"
)

type config struct {
	pokeApi        *pokeapi.PokeAPI
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {
	config := &config{
		pokeApi:        pokeapi.GetPokeAPI(),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for a list of available commands.")
	repl(config, scanner)
}

func repl(config *config, scanner *bufio.Scanner) {
	for {
		fmt.Println()
		fmt.Print("Pokedex> ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("There was an error reading your input. Please try again.")
			continue
		}

		inputWords := strings.Fields(strings.ToLower(scanner.Text()))
		if len(inputWords) == 0 {
			continue
		}

		command, ok := getCommands()[inputWords[0]]
		if !ok {
			fmt.Println("Invalid command. Use 'help' for a list of commands.")
			continue
		}

		err := command.callback(config, inputWords[1:])
		if err != nil {
			fmt.Println("There was an following error processing your command.", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Traverse locations forward.",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Traverse locations back.",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Lists all pokemons in given location.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Show pokemon information.",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp(_ *config, _ []string) error {
	fmt.Println("List of avalable commands:")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(_ *config, _ []string) error {
	fmt.Println("Thanks for using the Pokedex!")
	fmt.Println("Have a nice day!")
	os.Exit(0)
	return nil
}
