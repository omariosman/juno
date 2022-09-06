// Package gossip provides abstractions for pubsub based communication
// over the gossip protocol.
package gossip

import (
	"context"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
)

// Gossip represents a pubsub topic where messages are distributed using
// the gossip protocol.
type Gossip struct {
	topic        *pubsub.Topic
	subscription *pubsub.Subscription
}

// New creates a new Gossip.
func New(ctx context.Context, h host.Host, topic string) (*Gossip, error) {
	gs, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, err
	}

	handle, err := gs.Join(topic)
	if err != nil {
		return nil, err
	}

	subscription, err := handle.Subscribe()
	if err != nil {
		return nil, err
	}

	return &Gossip{topic: handle, subscription: subscription}, nil
}

// Publish publishes a message to a given topic.
func (g *Gossip) Publish(ctx context.Context, msg []byte) error {
	return g.topic.Publish(ctx, msg)
}

// Listen listens for messages on a given topic.
func (g *Gossip) Listen(ctx context.Context) error {
	for {
		msg, err := g.subscription.Next(ctx)
		if err != nil {
			return err
		}
		// TODO: Handle message received.
		fmt.Println("p2p/gossip: message received:", string(msg.Data))
	}
}
