package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://pokemon:pokemon@localhost:5432/pokemon?sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected succesfully !!")

	return conn, nil
}
