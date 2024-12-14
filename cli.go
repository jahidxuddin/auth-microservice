package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptForConfig() (string, string, error) {
	fmt.Print("Enter JWT Secret: ")
	jwtSecret, err := readInput()
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Data Source Name: ")
	dbHost, err := readInput()
	if err != nil {
		return "", "", err
	}

	return jwtSecret, dbHost, nil
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
