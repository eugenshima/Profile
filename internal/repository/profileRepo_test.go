package repository

import (
	"context"
	"testing"

	"github.com/eugenshima/profile/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var rps *ProfileRepository

var (
	testProfile = &model.Profile{
		ID:           uuid.New(),
		Login:        "test_login",
		Password:     "test_password",
		RefreshToken: "test_token",
	}
	testAuth = &model.Auth{
		Login:    "test_login",
		Password: "test_password",
	}
)

func TestGetIDByLoginPassword(t *testing.T) {
	err := CreateTestProfile()
	require.NoError(t, err)
	defer func() {
		err = DeleteTestProfile(testProfile.ID)
		require.NoError(t, err)
	}()

	id, pass, err := rps.GetIDByLoginPassword(context.Background(), testAuth.Login)
	require.NoError(t, err)
	require.NotNil(t, id)
	require.NotNil(t, pass)
	require.Equal(t, id, testProfile.ID)
}

func TestGetIdByWrongLoginPassword(t *testing.T) {
	err := CreateTestProfile()
	require.NoError(t, err)
	defer func() {
		err = DeleteTestProfile(testProfile.ID)
		require.NoError(t, err)
	}()

	id, pass, err := rps.GetIDByLoginPassword(context.Background(), "fake login")
	require.Error(t, err)
	require.Nil(t, id)
	require.Nil(t, pass)
}
