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
	callback func(*Config) error
}

type Config struct {
	Next string 
	Previous string 
	PageNo int 
}

var commands map[string]cliCommand

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	cfg := Config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

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
		"clear": {
			name: "clear",
			description: "Clear the entire screen",
			callback: commandClear,
		},
		"map": {
			name: "map",
			description: "Show the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show the previous 20 locations",
			callback: commandMapb,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		if reader.Text() == "" {
			continue
		}
		cleanText := cleanInput(reader.Text())
		userCommand := cleanText[0]

		if _, exists := commands[userCommand]; exists {
			commands[userCommand].callback(&cfg)
		} else {
			fmt.Println("Command not found! Try help more info")
		}

		fmt.Println()
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}

