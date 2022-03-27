package postgres

import (
	"context"
	"transbox/pkg/store"

	"github.com/jackc/pgx/v4"
)

type TxWrapper struct {
	tx pgx.Tx
}

func (w *TxWrapper) Exec(ctx context.Context, command func(ctx context.Context, q store.Querier) error) error {
	if err := command(ctx, w.tx); err != nil {
		_ = w.tx.Rollback(ctx)
		return err
	}

	return w.tx.Commit(ctx)
}
