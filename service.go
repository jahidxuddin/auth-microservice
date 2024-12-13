package main

func NewAuthService(db Database) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) register(creds Credentials) error {
	return nil
}

func (s *AuthService) login(creds Credentials) error {
	return nil
}
