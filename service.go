package main

import (
	"auth-microservice/ent"
	"auth-microservice/ent/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthService(db Database, jwtSecret string) *AuthService {
	return &AuthService{db: db, jwtSecret: jwtSecret}
}

func (s *AuthService) register(creds Credentials) error {
	_, err := s.db.instance.Client.User.Query().Where(user.Email(creds.Email)).Only(s.db.ctx)
	if err == nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.db.instance.Client.User.Create().SetEmail(creds.Email).SetPassword(string(hashedPassword)).Save(s.db.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) login(creds Credentials) (token string, failed error) {
	user, err := s.db.instance.Client.User.Query().Where(user.Email(creds.Email)).Only(s.db.ctx)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	secretKey := []byte(s.jwtSecret)
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Email,
		"admin": false,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err = generateJWT(secretKey, claims)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
}

func generateJWT(secretKey []byte, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *AuthService) verifyJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token: %v", err)
	}
	return nil
}

func (s *AuthService) getEmailFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["name"].(string)
		if !ok {
			return "", fmt.Errorf("email not found in token")
		}
		fmt.Println(email)
		return email, nil
	}

	return "", fmt.Errorf("invalid token claims")
}

func (s *AuthService) getUserFromToken(tokenString string) (*ent.User, error) {
	email, err := s.getEmailFromToken(tokenString)
	if err != nil {
		return nil, err
	}

	user, err := s.db.instance.Client.User.Query().Where(user.Email(email)).Only(s.db.ctx)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return user, nil
}

func (s *AuthService) updateUser(userID int, newEmail, newPassword string) (string, error) {
	update := s.db.instance.Client.User.UpdateOneID(userID)

	if newEmail != "" {
		update.SetEmail(newEmail)
	}

	if newPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		update.SetPassword(string(hashedPassword))
	}

	user, err := update.Save(s.db.ctx)
	if err != nil {
		return "nil", fmt.Errorf("failed to update user: %v", err)
	}

	secretKey := []byte(s.jwtSecret)
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Email,
		"admin": false,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err := generateJWT(secretKey, claims)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
}
