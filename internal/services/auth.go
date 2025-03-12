package services

import (
	"database/sql"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db        *sql.DB
	jwtSecret []byte // Add this to use a consistent secret
}

// NewAuthService initializes an AuthService with a database connection and JWT secret
func NewAuthService(db *sql.DB, jwtSecret []byte) *AuthService {
	if db == nil {
		panic("AuthService: db cannot be nil")
	}
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// Register creates a new user and returns their ID
func (s *AuthService) Register(email, password string) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	var userID int
	err = s.db.QueryRow(
		"INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", // Fixed typo: "passoword_hash" -> "password_hash"
		email,
		string(hash),
	).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Login authenticates a user and returns a JWT
func (s *AuthService) Login(email, password string) (string, error) {
	var userID int
	var passwordHash string
	err := s.db.QueryRow(
		"SELECT id, password_hash FROM users WHERE email = $1",
		email,
	).Scan(&userID, &passwordHash)
	if err == sql.ErrNoRows {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), // Added for completeness
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Changed ES256 to HS256
	tokenString, err := token.SignedString(s.jwtSecret)        // Use s.jwtSecret instead of config.LoadConfig()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}