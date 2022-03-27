package postgres

import (
	"context"
	"transbox/pkg/store"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Store implements repository.Store for postgres
type Store struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{db: pool}
}

func (s *Store) Connection(ctx context.Context) (store.Conn, error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	return &Conn{conn}, nil
}

func (s *Store) Tx(ctx context.Context) (store.TxWrapper, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &TxWrapper{
		tx: tx,
	}, nil
}
