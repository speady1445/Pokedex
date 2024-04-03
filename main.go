package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for a list of available commands.")
	repl(scanner, commands)
}

func repl(scanner *bufio.Scanner, commands map[string]cliCommand) {
	for {
		fmt.Println()
		fmt.Print("Pokedex> ")
		scanner.Scan()
		if scanner.Err() != nil {
			fmt.Println("There was an error reading your input. Please try again.")
			continue
		}

		command, ok := commands[strings.ToLower(scanner.Text())]
		if !ok {
			fmt.Println("Invalid command. Use 'help' for a list of commands.")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println("There was an error processing your command. Please try again.")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("List of avalable commands:")
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Thanks for using the Pokedex!")
	fmt.Println("Have a nice day!")
	os.Exit(0)
	return nil
}
