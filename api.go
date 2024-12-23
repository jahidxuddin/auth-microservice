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
	http.HandleFunc("/api/v1/update-user", LoggingMiddleware(s.updateUserHandler))
	http.HandleFunc("/api/v1/email", LoggingMiddleware(s.emailHandler))
	http.HandleFunc("/api/v1/find-user-by-jwt", LoggingMiddleware(s.findUserByJWTHandler))
	http.HandleFunc("/api/v1/verify-jwt", LoggingMiddleware(s.verifyJWTHandler))

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

	token, err := s.authService.login(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Login successful", "token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ApiServer) emailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization header is required", http.StatusBadRequest)
		return
	}

	email, err := s.authService.getEmailFromToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	response := map[string]string{"email": email}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ApiServer) findUserByJWTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization header is required", http.StatusBadRequest)
		return
	}

	user, err := s.authService.getUserFromToken(token)
	if err != nil {
		http.Error(w, "Invalid token or user not found", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ApiServer) verifyJWTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization header is required", http.StatusBadRequest)
		return
	}

	err := s.authService.verifyJWT(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	response := map[string]string{"message": "Token is valid"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *ApiServer) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization header is required", http.StatusBadRequest)
		return
	}

	user, err := s.authService.getUserFromToken(token)
	if err != nil {
		http.Error(w, "Invalid token or user not found", http.StatusUnauthorized)
		return
	}

	var updateData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err = s.authService.updateUser(user.ID, updateData.Email, updateData.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
