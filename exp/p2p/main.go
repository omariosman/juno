package main

import (
	"os"

	"gojuno.xyz/p2p/node"
)

func main() {
	// Initialise a node.
	n, err := node.New()
	if err != nil {
		panic(err)
	}

	// If an argument is provided then this is the node that will send
	// ping requests.
	if len(os.Args) > 1 {
		// Connect to node and send ping request.
		n.Ping(os.Args[1])
	}

	// DEBUG.
	// Hang and listen for connections.
	select {}
}
