package main

import "fmt"

func main() {
	port, jwtSecret, dataSourceName, err := PromptForConfig()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	database := NewDatabase(dataSourceName)
	authService := NewAuthService(*database, jwtSecret)
	apiServer := NewApiServer(port, *authService)
	apiServer.Start()
}
