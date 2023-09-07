// Package model of our entity
package model

import "github.com/google/uuid"

// Profile struct represents a Profile model
type Profile struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
