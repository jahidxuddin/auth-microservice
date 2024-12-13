package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewApiServer(port string, authService AuthService) *ApiServer {
	return &ApiServer{port: port, authService: authService}
}

func (s *ApiServer) Start() {
	http.HandleFunc("/api/v1/register", LoggingMiddleware(s.registerHandler))
	http.HandleFunc("/api/v1/login", LoggingMiddleware(s.loginHandler))

	fmt.Printf("Server is running on http://localhost:%s\n", s.port)
	err := http.ListenAndServe("localhost:"+s.port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func (s *ApiServer) registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	err = s.authService.register(creds)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ApiServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	err = s.authService.login(creds)
	if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

	response := map[string]string{"message": "Login successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
