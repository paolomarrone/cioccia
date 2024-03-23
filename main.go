package main

import (
	"fmt"
	"os/exec"
)

func executeCommand(command string, args []string, outputChan chan<- string) {
	cmd := exec.Command(command, args...)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		outputChan <- fmt.Sprintf("Error executing %s: %s", command, err)
		return
	}

	outputChan <- string(output)
}

func main() {
	commands := []struct {
		Command string
		Args    []string
	}{
		{"echo", []string{"Hello"}},
		{"ls", []string{"-l", "-a"}},
		{"date", []string{}},
	}

	outputChan := make(chan string)

	for _, cmd := range commands {
		go executeCommand(cmd.Command, cmd.Args, outputChan)
	}

	// Receive and print output from commands
	for i := 0; i < len(commands); i++ {
		output := <-outputChan
		fmt.Println(output)
	}
}
