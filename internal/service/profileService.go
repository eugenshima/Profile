// Package service provides a set of functions, which include business-logic in it
package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/profile/internal/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ProfileService struct represents a profile service
type ProfileService struct {
	rps ProfileRepositoryInterface
}

// NewProfileService creates a new ProfileService
func NewProfileService(rps ProfileRepositoryInterface) *ProfileService {
	return &ProfileService{rps: rps}
}

// ProfileRepositoryInterface represents a profile repository methods
type ProfileRepositoryInterface interface {
	GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, profile *model.Profile) error
	GetIDByLoginPassword(ctx context.Context, login string) (uuid.UUID, []byte, error)
	DeleteProfileByID(ctx context.Context, id uuid.UUID) error
}

// GetProfileByID returns a profile by given ID
func (s *ProfileService) GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	return s.rps.GetProfileByID(ctx, id)
}

// CreateNewProfile function creates new profile
func (s *ProfileService) CreateNewProfile(ctx context.Context, profile *model.Profile) error {
	return s.rps.CreateProfile(ctx, profile)
}

func (s *ProfileService) UpdateProfile(ctx context.Context, profile *model.Profile) error {
	return s.rps.UpdateProfile(ctx, profile)
}

func (s *ProfileService) Login(ctx context.Context, login *model.Auth) (uuid.UUID, error) {
	id, password, err := s.rps.GetIDByLoginPassword(ctx, login.Login)
	if err != nil {
		return uuid.Nil, fmt.Errorf("GetIDByLoginPassword: %w", err)
	}
	err = bcrypt.CompareHashAndPassword(password, login.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("wrong password: %w", err)
	}
	return id, nil
}

func (s *ProfileService) DeleteProfileByID(ctx context.Context, id uuid.UUID) error {
	return s.rps.DeleteProfileByID(ctx, id)
}
