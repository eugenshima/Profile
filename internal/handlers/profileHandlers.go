package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/profile/internal/model"
	proto "github.com/eugenshima/profile/proto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProfileHandler struct {
	srv ProfileServiceInterface
	proto.UnimplementedPriceServiceServer
}

func NewProfileHandler(srv ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{srv: srv}
}

type ProfileServiceInterface interface {
	GetProfileByID(ctx context.Context, id uuid.UUID) (*model.Profile, error)
	CreateNewProfile(ctx context.Context, profile *model.Profile) error
}

func (ph *ProfileHandler) GetProfileByID(ctx context.Context, req *proto.GetProfileByIDRequest) (*proto.GetProfileByIDResponse, error) {
	ID, err := uuid.Parse(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"req.User.Balance_ID": req.ID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	profile, err := ph.srv.GetProfileByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ctx": ctx, "ID": ID}).Errorf("GetProfileByID: %v", err)
		return nil, fmt.Errorf("GetProfileByID: %w", err)
	}
	protoProfile := &proto.Profile{
		ID:       profile.ID.String(),
		Login:    profile.Login,
		Password: profile.Password,
	}
	return &proto.GetProfileByIDResponse{Profile: protoProfile}, nil
}

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

func (ph *ProfileHandler) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	return &proto.UpdateProfileResponse{}, nil
}
