package database

import (
	"context"
	"fmt"
	"pokemon-platform/backend/config"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(
		context.Background(),
		config.App.DatabaseURL,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected succesfully !!")

	return conn, nil
}
