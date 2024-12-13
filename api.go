package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewApiServer(port string) *ApiServer {
	return &ApiServer{port: port}
}

func (s *ApiServer) Start() {
	http.HandleFunc("/api/v1/register", LoggingMiddleware(registerHandler))
	http.HandleFunc("/api/v1/login", LoggingMiddleware(loginHandler))

	fmt.Printf("Server is running on http://localhost:%s\n", s.port)
	err := http.ListenAndServe("localhost:"+s.port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"message": "Login successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
