package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands map[string]cliCommand

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	commands = map[string]cliCommand{
		"help": {
			name: "help",
			description: "Display this help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()


		cleanText := cleanInput(reader.Text())
		userCommand := cleanText[0]

		if _, exists := commands[userCommand]; exists {
			if err := commands[userCommand].callback(); err != nil {
				errMessage := fmt.Errorf("error occured during running the command: %w", err)
				fmt.Println(errMessage)
			}
		} else {
			fmt.Println("Command not found! Enter help more info")
		}
		fmt.Println()
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}

