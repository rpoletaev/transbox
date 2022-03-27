package service

import (
	"context"
	"strconv"
	"transbox/internal/outbox"
)

func (s Service) Create(ctx context.Context, topic string, entityID uint) error {
	event := outbox.Event{
		Topic:  topic,
		Data:   []byte(strconv.Itoa(int(entityID))),
		Status: string(outbox.EventStatusCreated),
	}
	return s.events.Create(ctx, event)
}
