// Package model of our entity
package model

import "github.com/google/uuid"

// Profile struct represents a Profile model
type Profile struct {
	ID           uuid.UUID `json:"id"`
	BalanceID    uuid.UUID `json:"balance_id"`
	Login        string    `json:"login"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}
