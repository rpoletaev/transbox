package repository

import (
	"context"
	"transbox/internal/outbox"
	"transbox/pkg/store"
)

type EventsRepository struct {
	q store.Querier
}

func New(q store.Querier) *EventsRepository {
	return &EventsRepository{q: q}
}

func (r *EventsRepository) Create(ctx context.Context, e outbox.Event) error {
	_, err := r.q.Exec(ctx, "INSERT INTO events (topic, data, status) VALUES($1, $2, $3)", e.Topic, e.Data, e.Status)
	return err
}
