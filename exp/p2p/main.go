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
		if err := n.Ping(os.Args[1]); err != nil {
			panic("failed to execute ping: " + err.Error())
		}
	}

	// DEBUG.
	// Hang and listen for connections.
	select {}
}
