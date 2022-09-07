// Binary p2p initialises a node that will connect to the IPFS bootstrap
// nodes and subscribe to the topic starknet_test_nodes.
//
// To run the application, execute the following command.
//
//	go run .
//
// And then run the following in another terminal instance to launch the
// peer that will execute a ping request after it has discovered another
// peer subscribed to starknet_test_nodes that also follows the
// starknet/0.10.0 protocol.
//
//	go run . -peer
//
// Both nodes will then publish pubsub messages identifying themselves
// every predefined interval.
package main

import (
	"context"
	"flag"
	"time"

	"gojuno.xyz/p2p/gossip"
	"gojuno.xyz/p2p/node"
	"gojuno.xyz/p2p/protocol/ping"
)

func main() {
	peer := flag.Bool("peer", false, "init secondary node that will execute ping")
	flag.Parse()

	ctx := context.Background()

	// Initialise a node.
	n, err := node.New(ctx)
	if err != nil {
		panic("p2p/main: new node: " + err.Error())
	}

	/* Register protocols. */

	// Register custom ping protocol.
	pp := ping.Register(n.Host)

	// TODO: The discovery topic should be configured as
	// "_starknet_nodes/" + configured chain id. See the following for
	// details about how chain ids are encoded https://docs.starknet.io/docs/Blocks/transactions/#chain-id.
	const topic = "starknet_test_nodes"
	gs, err := gossip.New(ctx, n.Host, topic)
	if err != nil {
		panic("p2p/main: new gossip: " + err.Error())
	}

	/* Conduct peer discovery. */

	// Discover a peer. It is important that all protocols are registered
	// and configured before peer discovery is conducted (lines preceding
	// this comment).
	pi, err := n.Discover(ctx, topic)
	if err != nil {
		panic("p2p/main: discover: " + err.Error())
	}

	/* Execute protocols. */

	// If an argument is provided then this is the node that will send a
	// ping request "PING" while the other will reply with "PONG".
	if *peer {
		if err := pp.Ping(ctx, pi); err != nil {
			panic("p2p/main: ping: " + err.Error())
		}
	}

	// Listen on the topic subscribed.
	go func() {
		if err := gs.Listen(ctx); err != nil {
			panic("p2p/main: listen: " + err.Error())
		}
	}()

	// Message that will identify a peer publishing the message on the
	// topic.
	var which string
	if *peer {
		which = "two"
	} else {
		which = "one"
	}

	// Publish messages on the topic at every interval.
	const interval = 5 * time.Second
	for {
		if err := gs.Publish(ctx, []byte("node: "+which)); err != nil {
			panic("p2p/main: publish: " + err.Error())
		}
		time.Sleep(interval)
	}

	// DEBUG.
	// Hang and listen for connections. Doing a graceful shutdown would
	// probably be better.
	// select {}
}
