package repository

import (
	"context"
	"transbox/internal/entities"
	"transbox/pkg/store"
)

type EntityRepository struct {
	q store.Querier
}

func New(q store.Querier) *EntityRepository {
	return &EntityRepository{q: q}
}

func (r *EntityRepository) Create(ctx context.Context, e entities.Entity) error {
	_, err := r.q.Exec(ctx, "insert into entities (name) values($1)", e.ID, e.Name)
	return err
}
