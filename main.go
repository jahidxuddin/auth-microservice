package main

func main() {
	database := NewDatabase("root:admin123@tcp(localhost:3306)/authdb?parseTime=True")
	authService := NewAuthService(*database)
	apiServer := NewApiServer("8080", *authService)
	apiServer.Start()
}
