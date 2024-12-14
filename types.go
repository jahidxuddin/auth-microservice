package main

import "context"

type Database struct {
	instance *DBClient
	ctx      context.Context
}

type AuthService struct {
	db Database
	jwtSecret string
}

type ApiServer struct {
	port        string
	authService AuthService
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
