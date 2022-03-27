package service

import (
	"context"
	"transbox/internal/outbox"
)

type EventsRepository interface {
	Create(ctx context.Context, e outbox.Event) error
}

type EventProducer interface {
	Send(ctx context.Context, e outbox.Event) error
}

type Service struct {
	events   EventsRepository
	producer EventProducer
}

func New(r EventsRepository, p EventProducer) Service {
	return Service{
		events:   r,
		producer: p,
	}
}
