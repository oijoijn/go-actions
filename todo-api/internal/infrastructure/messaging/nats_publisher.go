package messaging

import (
	"context"
	"todo-app/internal/usecase"

	"github.com/nats-io/nats.go"
)

type natsPublisher struct {
	js nats.JetStreamContext
}

func NewNatsPublisher(js nats.JetStreamContext) usecase.EventPublisher {
	return &natsPublisher{js: js}
}

func (p *natsPublisher) Publish(ctx context.Context, subject string, data []byte) error {
	_, err := p.js.Publish(subject, data)
	return err
}
