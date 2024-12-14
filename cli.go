package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func PromptForConfig() (string, string, string, error) {
	fmt.Print("Enter Port: ")
	port, err := readInput()
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Enter JWT Secret: ")
	jwtSecret, err := readInput()
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Enter Data Source Name: ")
	dbHost, err := readInput()
	if err != nil {
		return "", "", "", err
	}

	clearTerminal()

	return port, jwtSecret, dbHost, nil
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func clearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear") 
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
