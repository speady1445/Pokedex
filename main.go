package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/speady1445/Pokedex/internals/pokeapi"
)

type config struct {
	pokeApi *pokeapi.PokeAPI
}

func main() {
	config := &config{
		pokeApi: pokeapi.GetPokeAPI(),
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

		command, ok := getCommands()[strings.ToLower(scanner.Text())]
		if !ok {
			fmt.Println("Invalid command. Use 'help' for a list of commands.")
			continue
		}

		err := command.callback(config)
		if err != nil {
			fmt.Println("There was an following error processing your command.", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp(_ *config) error {
	fmt.Println("List of avalable commands:")
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandExit(_ *config) error {
	fmt.Println("Thanks for using the Pokedex!")
	fmt.Println("Have a nice day!")
	os.Exit(0)
	return nil
}
