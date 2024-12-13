package main

import "context"

type Database struct {
	dbClient *DBClient
	ctx      context.Context
}

type AuthService struct {
	db Database
}

type ApiServer struct {
	port        string
	authService AuthService
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
