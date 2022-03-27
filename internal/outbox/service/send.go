package service

import (
	"context"
	"transbox/internal/outbox"
)

func (s Service) Send(ctx context.Context, e outbox.Event) error {

	return nil
}
