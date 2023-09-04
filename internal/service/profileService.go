package service

import (
	"context"

	"github.com/eugenshima/profile/internal/model"

	"github.com/google/uuid"
)

type ProfileService struct {
	rps ProfileRepositoryInterface
}

func NewProfileService(rps ProfileRepositoryInterface) *ProfileService {
	return &ProfileService{rps: rps}
}

type ProfileRepositoryInterface interface {
	GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error)
	CreateProfile(ctx context.Context, profile *model.Profile) error
}

func (s *ProfileService) GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error) {
	return s.rps.GetProfileByID(ctx, id)
}

func (s *ProfileService) CreateNewProfile(ctx context.Context, profile *model.Profile) error {
	return s.rps.CreateProfile(ctx, profile)
}
