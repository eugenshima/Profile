package handlers

import (
	"context"
	"os"
	"testing"

	"github.com/eugenshima/profile/internal/handlers/mocks"
	"github.com/eugenshima/profile/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	mockProfileService *mocks.ProfileService
	mockAuth           = &model.Auth{
		Login:    "test_login",
		Password: "test_password",
	}
)

// TestMain execute all tests
func TestMain(m *testing.M) {
	mockProfileService = new(mocks.ProfileService)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestHandlerLogin(t *testing.T) {
	mockProfileService.On("Login", mock.Anything, mock.AnythingOfType("*model.Auth")).Return(uuid.UUID{}, nil).Once()

	id, err := mockProfileService.Login(context.Background(), mockAuth)
	require.NoError(t, err)
	require.NotNil(t, id)

	assertion := mockProfileService.AssertExpectations(t)
	require.True(t, assertion)
}

func TestGetProfileByID(t *testing.T) {
	mockProfileService.On("GetProfileByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&model.Profile{}, nil).Once()

	ID := uuid.New()
	id, err := mockProfileService.GetProfileByID(context.Background(), ID)
	require.NoError(t, err)
	require.NotNil(t, id)

	assertion := mockProfileService.AssertExpectations(t)
	require.True(t, assertion)
}

func TestCreateNewProfile(t *testing.T) {
	//TODO: :)
}
