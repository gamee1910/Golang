package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		//Read the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		//Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	//Remove the new line character
	input = strings.TrimSuffix(input, "\n")

	//Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	//Check for built-in commands
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supproted
		if len(args) < 2 {
			return errors.New("Path required")
		}

		//Change the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	//Pass the program and the arguments spearately
	cmd := exec.Command(args[0], args[1:]...)

	//Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//Execute the command
	return cmd.Run()
}
