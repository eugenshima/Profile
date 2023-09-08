// Package model of our entity
package model

import "github.com/google/uuid"

// Profile struct represents a Profile model
type Profile struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	Password     []byte    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
	Username     string    `json:"username"`
}

type Auth struct {
	Login    string `json:"login"`
	Password []byte `json:"password"`
}
