package eventbus

import "context"

type Bus interface {
	Send(ctx context.Context, topic string) error
}
