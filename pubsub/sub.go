package pubsub

import (
	players "club-service/client"
	"context"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

// All methods of Sub struct will be executed when a message is received
type Sub struct{}

func (s *Sub) Process(ctx context.Context, event *players.Player) error {
	md, _ := metadata.FromContext(ctx)
	logger.Infof("[pubsub.1] Received event %+v with metadata %+v", event, md)
	// do something with event
	return nil
}
