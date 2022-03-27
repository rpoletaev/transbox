package store

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Conn interface {
	Querier
	Close()
}

type Store interface {
	Connection(context.Context) (Conn, error)
	Tx(context.Context) (TxWrapper, error)
}

type TxWrapper interface {
	Exec(ctx context.Context, command func(ctx context.Context, q Querier) error) error
}
