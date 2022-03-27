package main

import (
	"context"
	"transbox/internal/usecases"
	"transbox/pkg/store/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	pool, err := pgxpool.Connect(context.Background(), "postgres")
	if err != nil {
		panic(err)
	}

	pgStore := postgres.New(pool)
	createEntityUsecase := usecases.NewCreateEntity(pgStore)
}
