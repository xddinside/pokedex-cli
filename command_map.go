package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *Config) error {
	resp, err := http.Get(cfg.Next)
	if err != nil {
		fmt.Println("Can't go any further ahead!'")
		return err
	}
	defer resp.Body.Close()

	cfg.PageNo++
	fmt.Printf("\nPage: %d", cfg.PageNo)
	
	var payload JsonData
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&payload); err != nil {
		return err
	}

	cfg.Next = payload.Next
	cfg.Previous = payload.Previous
	for _, area := range payload.Results {
		fmt.Printf("\n	%s", area.Name)
	}
	fmt.Println()

	return nil
}
