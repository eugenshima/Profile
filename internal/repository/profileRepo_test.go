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
		Password:     []byte("test_password"),
		RefreshToken: []byte("test_token"),
	}
	testAuth = &model.Auth{
		Login:    "test_login",
		Password: []byte("test_password"),
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
	require.NoError(t, err)
	require.Equal(t, id, uuid.Nil)
	require.Equal(t, pass, "")
}

func TestGetProfileByID(t *testing.T) {
	err := CreateTestProfile()
	require.NoError(t, err)
	defer func() {
		err = DeleteTestProfile(testProfile.ID)
		require.NoError(t, err)
	}()

	profile, err := rps.GetProfileByID(context.Background(), testProfile.ID)
	require.NoError(t, err)
	require.NotNil(t, profile)
	require.Equal(t, profile.ID, testProfile.ID)
	require.Equal(t, profile.Login, testProfile.Login)
	require.Equal(t, profile.Password, testProfile.Password)
}

func TestGetProfileByWrongID(t *testing.T) {
	err := CreateTestProfile()
	require.NoError(t, err)
	defer func() {
		err = DeleteTestProfile(testProfile.ID)
		require.NoError(t, err)
	}()

	profile, err := rps.GetProfileByID(context.Background(), uuid.New())
	require.Error(t, err)
	require.Nil(t, profile)
}

func TestCreateProfile(t *testing.T) {
	err := CreateTestProfile()
	require.NoError(t, err)
	defer func() {
		err = DeleteTestProfile(testProfile.ID)
		require.NoError(t, err)
	}()
}
