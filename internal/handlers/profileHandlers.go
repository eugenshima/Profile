// Package handlers contains gRPC methods
package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/profile/internal/model"
	proto "github.com/eugenshima/profile/proto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// ProfileHandler struct represents profile handler
type ProfileHandler struct {
	srv ProfileServiceInterface
	proto.UnimplementedPriceServiceServer
}

// NewProfileHandler function creates a new profile handler
func NewProfileHandler(srv ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{srv: srv}
}

// ProfileServiceInterface interface represents service interface methods
type ProfileServiceInterface interface {
	GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error)
	CreateNewProfile(ctx context.Context, profile *model.Profile) error
	UpdateProfile(ctx context.Context, profile *model.Profile) error
	Login(ctx context.Context, loginPass *model.Auth) (uuid.UUID, error)
}

func (ph *ProfileHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	infoToLogin := &model.Auth{
		Login:    req.Auth.Login,
		Password: req.Auth.Password,
	}
	ID, err := ph.srv.Login(ctx, infoToLogin)
	if err != nil {
		logrus.WithFields(logrus.Fields{"infoToLogin": infoToLogin}).Errorf("Login: %v", err)
		return nil, fmt.Errorf("Login: %w", err)
	}
	return &proto.LoginResponse{ID: ID.String()}, nil
}

// GetProfileByID function gets profile vy provided ID
func (ph *ProfileHandler) GetProfileByID(ctx context.Context, req *proto.GetProfileByIDRequest) (*proto.GetProfileByIDResponse, error) {
	ID, err := uuid.Parse(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": req.ID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	profile, err := ph.srv.GetProfileByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ctx": ctx, "ID": ID}).Errorf("GetProfileByID: %v", err)
		return nil, fmt.Errorf("GetProfileByID: %w", err)
	}
	protoProfile := &proto.Profile{
		ID:           profile.ID.String(),
		Login:        profile.Login,
		Password:     profile.Password,
		RefreshToken: profile.RefreshToken,
	}
	return &proto.GetProfileByIDResponse{Profile: protoProfile}, nil
}

// CreateNewProfile function creates a new profile
func (ph *ProfileHandler) CreateNewProfile(ctx context.Context, req *proto.CreateNewProfileRequest) (*proto.CreateNewProfileResponse, error) {
	newProfile := &model.Profile{
		ID:       uuid.New(),
		Login:    req.Auth.Login,
		Password: req.Auth.Password,
	}
	err := ph.srv.CreateNewProfile(ctx, newProfile)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ctx": ctx, "newProfile": newProfile}).Errorf("CreateNewProfile: %v", err)
		return nil, fmt.Errorf("CreateNewProfile: %w", err)
	}
	return &proto.CreateNewProfileResponse{}, nil
}

// UpdateProfile function updates a profile information
func (ph *ProfileHandler) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	ID, err := uuid.Parse(req.Profile.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": req.Profile.ID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	ProfileToUpdate := &model.Profile{
		ID:           ID,
		Login:        req.Profile.Login,
		Password:     req.Profile.Password,
		RefreshToken: req.Profile.RefreshToken,
	}
	err = ph.srv.UpdateProfile(ctx, ProfileToUpdate)
	if err != nil {
		logrus.WithFields(logrus.Fields{"newProfile": ProfileToUpdate}).Errorf("UpdateProfile: %v", err)
		return nil, fmt.Errorf("UpdateProfile: %w", err)
	}
	return &proto.UpdateProfileResponse{}, nil
}
