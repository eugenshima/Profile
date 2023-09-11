// Package model of our entity
package model

import "github.com/google/uuid"

// Profile struct represents a Profile model
type Profile struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	Password     []byte    `json:"password"`
	RefreshToken []byte    `json:"refresh_token"`
	Username     string    `json:"username"`
}

type Auth struct {
	Login    string `json:"login"`
	Password []byte `json:"password"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateTokens struct {
	ID           uuid.UUID `json:"id"`
	RefreshToken []byte    `json:"refresh_token"`
}
