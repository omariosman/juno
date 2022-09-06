// Binary p2p initialises a node that will connect to the IPFS bootstrap
// nodes and subscribe to the topic starknet_test_nodes.
//
// To run the application, execute the following command.
//
//	go run .
//
// And then run the following in another terminal instance to run the
// peer that will execute a ping request after it has discovered another
// peer subscribed to starknet_test_nodes that also follows the
// starknet/0.10.0 protocol.
//
//	go run . -peer
package main

import (
	"context"
	"flag"

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

	// Register ping protocol.
	pp := ping.Register(n.Host)

	// Discover a peer.
	// TODO: The discovery topic should be configured as
	// "_starknet_nodes/" + configured chain id. See the following for
	// details about how chain ids are encoded
	// https://docs.starknet.io/docs/Blocks/transactions/#chain-id.
	const topic = "starknet_test_nodes"
	pi, err := n.Discover(ctx, topic)
	if err != nil {
		panic("p2p/main: discover: " + err.Error())
	}

	// If an argument is provided then this is the node that will send
	// ping requests.
	if *peer {
		// Connect to node and send ping request.
		if err := pp.Ping(ctx, pi); err != nil {
			panic("p2p/main: ping: " + err.Error())
		}
	}

	// DEBUG.
	// Hang and listen for connections. Doing a graceful shutdown would
	// probably be better.
	select {}
}
