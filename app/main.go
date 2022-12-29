package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	var exitErr *exec.ExitError

	if err != nil {
		if errors.As(err, &exitErr) {
			fmt.Printf("program exited with status: %d\n", exitErr.ExitCode())
			os.Exit(exitErr.ExitCode())
		}

		fmt.Printf("Err: %s\n", err.Error())
		os.Exit(1)
	}

}
