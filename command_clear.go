package main

import (
	"os"
	"os/exec"
	"runtime"
)

func commandClear(cfg *Config) error {
	userOS := runtime.GOOS
	var cmd *exec.Cmd
	if userOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}
