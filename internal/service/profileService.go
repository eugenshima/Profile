// Package service provides a set of functions, which include business-logic in it
package service

import (
	"context"

	"github.com/eugenshima/profile/internal/model"

	"github.com/google/uuid"
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
}

// GetProfileByID returns a profile by given ID
func (s *ProfileService) GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	return s.rps.GetProfileByID(ctx, id)
}

// CreateNewProfile function creates new profile
func (s *ProfileService) CreateNewProfile(ctx context.Context, profile *model.Profile) error {
	return s.rps.CreateProfile(ctx, profile)
}
