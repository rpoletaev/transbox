package producer

import (
	"context"
	"transbox/internal/outbox"
	"transbox/pkg/eventbus"
)

type Producer struct {
	bus eventbus.Bus
}

func New(b eventbus.Bus) Producer {
	return Producer{bus: b}
}

func (p Producer) Send(ctx context.Context, e outbox.Event) error {
	return p.bus.Send(ctx, e.Topic)
}
