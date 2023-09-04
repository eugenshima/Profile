package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/profile/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type ProfileRepository struct {
	pool *pgxpool.Pool
}

func NewProfileRepository(pool *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{pool: pool}
}

func (db *ProfileRepository) GetAllProfiles(ctx context.Context) ([]*model.Profile, error) {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	users := []*model.Profile{}
	rows, err := db.pool.Query(ctx, "SELECT id, login, password FROM profile.profile")
	for rows.Next() {
		user := &model.Profile{}
		err := rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("Scan(): %w", err) // Returning error message
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *ProfileRepository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	tx, err := db.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	_, err = db.pool.Exec(ctx, "INSERT INTO profile.profile (id, login, password) VALUES ($1, $2, $3)", profile.ID, profile.Login, profile.Password)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
