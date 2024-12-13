package main

func main() {
	apiServer := NewApiServer("8080")
	apiServer.Start()
}
