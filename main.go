package main

import (
	"context"
	"fmt"

	cfgrtn "github.com/eugenshima/profile/internal/config"
	"github.com/eugenshima/profile/internal/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewDBPsql function provides Connection with PostgreSQL database
func NewDBPsql(env string) (*pgxpool.Pool, error) {
	// Initialization a connect configuration for a PostgreSQL using pgx driver
	config, err := pgxpool.ParseConfig(env)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}

	// Establishing a new connection to a PostgreSQL database using the pgx driver
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}
	// Output to console
	fmt.Println("Connected to PostgreSQL!")

	return pool, nil
}

func main() {
	cfg, err := cfgrtn.NewConfig()
	if err != nil {
		fmt.Printf("Error extracting env variables: %v", err)
		return
	}
	pool, err := NewDBPsql(cfg.PgxDBAddr)
	if err != nil {
		logrus.WithFields(logrus.Fields{"PgxDBAddr: ": cfg.PgxDBAddr}).Errorf("NewDBPsql: %v", err)
	}
	// example := &model.Profile{
	// 	ID:       uuid.New(),
	// 	Login:    "eugen",
	// 	Password: "password",
	// }
	repository := repository.NewProfileRepository(pool)
	//repository.CreateProfile(context.Background(), example)
	result, _ := repository.GetAllProfiles(context.Background())
	for _, value := range result {
		value.BalanceID = uuid.New()
		value.RefreshToken = "refresh token"
		fmt.Println(value)
	}
}
