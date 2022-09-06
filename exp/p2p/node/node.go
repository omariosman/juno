package node

import (
	"context"
	"errors"
	"fmt"
	"sync"

	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	discoveryutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/libp2p/go-libp2p/p2p/protocol/identify"
	"gojuno.xyz/p2p/gossip"
	"gojuno.xyz/p2p/protocol/ping"
)

// XXX: The user agent string can probably be generated at compile time
// using linker flags by reading the most recent Git tag.

const (
	userAgent       = "juno/0.3.0"
	protocolVersion = "starknet/0.10.0"
)

// ErrProtocolMismatch represents an error where the peer does not
// support the StarkNet protocol.
var ErrProtocolMismatch = errors.New("p2p/node: protocol mismatch")

// KeyGenError records an error that occurred during key-pair
// generation.
type KeyGenError struct {
	Err error
}

func (e *KeyGenError) Error() string {
	return fmt.Sprintf("p2p/node: key generation: %v", e.Err)
}

func (e *KeyGenError) Unwrap() error {
	return e.Err
}

// InitError records an error that occurred during initialisation of a
// new peer-to-peer node.
type InitError struct {
	Err error
}

func (e *InitError) Error() string {
	return fmt.Sprintf("p2p/node: node configuration: %v", e.Err)
}

func (e *InitError) Unwrap() error {
	return e.Err
}

// Node represents a node on the StarkNet peer-to-peer network.
type Node struct {
	host host.Host
	gs   *gossip.Gossip
	pp   *ping.Protocol
}

// New creates a new StarkNet node.
func New(ctx context.Context) (*Node, error) {
	// DEBUG.
	// While the specification requires that only Ed25519 keys are used,
	// the mainline IPFS DHT and default IPFS bootstrap nodes that are
	// used in the current discovery protocol use RSA keys so for now skip
	// key configuration.
	// pvt, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1 /* not applicable */)
	// if err != nil {
	// 	return nil, &KeyGenError{err}
	// }

	// Configure the host using parameters in the specification. For
	// everything else, use defaults.
	host, err := libp2p.New(
		// DEBUG.
		// See above. Also avoid configuring the network for now.
		/*
			libp2p.Identity(pvt),
			libp2p.ListenAddrStrings(
				"/ip4/0.0.0.0/tcp/0",
				"/ip6/::/tcp/0",
			),
			libp2p.Transport(tcp.NewTCPTransport),
		*/
		libp2p.UserAgent(userAgent),
		libp2p.ProtocolVersion(protocolVersion),
		libp2p.Ping(false),
	)
	if err != nil {
		return nil, &InitError{err}
	}

	// /p2p/id/delta/1.0.0 does not appear to be part of the
	// specification. See the following for details
	// https://github.com/libp2p/specs/blob/master/identify/README.md.
	host.RemoveStreamHandler(identify.IDDelta)

	// DEBUG.
	// Encode address information to multiaddr and print them to standard
	// output for debugging purposes.
	ai := peer.AddrInfo{ID: host.ID(), Addrs: host.Addrs()}
	addrs, err := peer.AddrInfoToP2pAddrs(&ai)
	if err != nil {
		return nil, err
	}
	for _, ma := range addrs {
		fmt.Println(ma)
	}

	// DEBUG.
	// The specification requires that nodes are able to connect using CID
	// encoded multihashes. libp2p plans to move towards this in future
	// versions but at time of writing both Go and Rust libraries still
	// default to base58 encoding. On the wire communication is done over
	// bytes so this distinction is largely immaterial.
	// println("p2p/node: cid encoded address: peer.ToCid(host.ID()).String())

	pp := ping.Register(host)

	return &Node{host: host, pp: pp}, nil
}

// connect connects to a peer with address information pi and returns an
// error otherwise.
func (n *Node) connect(ctx context.Context, pi peer.AddrInfo) error {
	if err := n.host.Connect(ctx, pi); err != nil {
		return err
	}

	// DEBUG.
	// Assert that the peer follows the StarkNet protocol. A better
	// implementation of this might use semantic versioning to determine
	// the appropriate level of compatibility.
	v, err := n.host.Peerstore().Get(pi.ID, "ProtocolVersion")
	if err != nil {
		return err
	}
	if v != protocolVersion {
		return ErrProtocolMismatch
	}

	return nil
}

// newTable creates a new distributed hash table and connects to one of
// the IPFS bootstrap nodes.
func (n *Node) newTable(ctx context.Context) (*dht.IpfsDHT, error) {
	table, err := dht.New(ctx, n.host)
	if err != nil {
		return nil, err
	}

	if err := table.Bootstrap(ctx); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	for _, addr := range dht.GetDefaultBootstrapPeerAddrInfos() {
		wg.Add(1)
		go func(pi peer.AddrInfo) {
			defer wg.Done()
			if err := n.host.Connect(ctx, pi); err != nil {
				// DEBUG.
				// There are multiple bootstrap nodes that one will attempt to
				// connect to so it should not be an issue if one of those
				// fails.
				fmt.Println("p2p/node: bootstrap connection:", err)
			}
		}(addr)
	}
	wg.Wait()

	return table, nil
}

// Discover finds peers to connect to that are subscribed to a given
// (discovery) topic.
func (n *Node) Discover(ctx context.Context, topic string) (peer.ID, error) {
	table, err := n.newTable(ctx)
	if err != nil {
		return "", err
	}

	discovery := routing.NewRoutingDiscovery(table)
	discoveryutil.Advertise(ctx, discovery, topic)

	for {
		ch, err := discovery.FindPeers(ctx, topic)
		if err != nil {
			return "", err
		}

		for peer := range ch {
			// Ignore ids that reference this node.
			if peer.ID == n.host.ID() {
				continue
			}

			if err := n.connect(ctx, peer); err != nil {
				// DEBUG.
				// There are multiple peers the node will attempt to connect to
				// and it is fine if either one of those attempts fail as long
				// as one succeeds.
				fmt.Println("p2p/node: peer connection:", err)
				continue
			}

			// DEBUG.
			fmt.Println("p2p/node: connected:", peer.ID.String())
			return peer.ID, nil
		}
	}
}

// Ping sends a ping request to the node with peerId pi.
// NOTE: This is not the standard libp2p ping protocol.
func (node *Node) Ping(ctx context.Context, pi peer.ID) error {
	return node.pp.Ping(ctx, pi)
}
