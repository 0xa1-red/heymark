package model

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound   error = fmt.Errorf("User not found")
	ErrUsernameExists error = fmt.Errorf("Username already exists")
	ErrUnauthorized   error = fmt.Errorf("Unauthorized")
)

type UserRepository interface {
	Authenticate(username, password string) (Token, error)
	Authorize(auth AuthData) (User, error)
	GetUser(id uuid.UUID) (User, error)
	CreateUser(username, password, bio string) (User, error)
}

type Token struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ValidUntil time.Time `json:"valid_until"`
}

func NewToken() Token {
	return Token{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		ValidUntil: time.Now().Add(1 * time.Hour),
	}
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Bio      string    `json:"bio"`
	Password string    `json:"-"`
	Tokens   []Token   `json:"-"`
}

type AuthData struct {
	Username string    `json:"username"`
	TokenID  uuid.UUID `json:"token"`
}

func GetAuthDataFromHeader(header string) (AuthData, error) {
	if header == "" {
		return AuthData{}, nil
	}
	encoded := strings.TrimPrefix(header, "Bearer ")
	raw, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return AuthData{}, fmt.Errorf("Error decoding base64 string: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(raw))
	var a AuthData
	if err := decoder.Decode(&a); err != nil {
		return AuthData{}, fmt.Errorf("Error decoding JSON: %w", err)
	}

	return a, nil
}
