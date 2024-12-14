package main

import "fmt"

func main() {
	jwtSecret, dataSourceName, err := PromptForConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	database := NewDatabase(dataSourceName)
	authService := NewAuthService(*database, jwtSecret)
	apiServer := NewApiServer("8080", *authService)
	apiServer.Start()
}
