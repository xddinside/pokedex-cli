package main 

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, command := range commands {
		fmt.Printf("	%s: %s\n", command.name, command.description)
	}
	return nil
}
