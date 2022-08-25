package node

import (
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	multiaddr "github.com/multiformats/go-multiaddr"
	"gojuno.xyz/p2p/protocol/ping"
)

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
	pp   *ping.Protocol
}

// New creates a new StarkNet node and prints its addresses to standard
// output. This is a temporary measure in the absence of a discovery
// protocol.
func New() (*Node, error) {
	pvt, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1 /* not applicable */)
	if err != nil {
		return nil, &KeyGenError{err}
	}

	// Configure the host using parameters in the specification. For
	// everything else, use defaults.
	host, err := libp2p.New(
		libp2p.Identity(pvt),
		libp2p.ListenAddrStrings(
			"/ip4/0.0.0.0/tcp/0",
			"/ip6/::/tcp/0",
		),
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.UserAgent("juno/0.3.0"),
	)
	if err != nil {
		return nil, &InitError{err}
	}

	pp := ping.Register(host)

	// DEBUG.
	// Encode address information to multiaddr and print them to standard
	// output so that that the peer can be setup to connect to this node
	// using said addresses in the absence of a discovery mechanism.
	ai := peer.AddrInfo{ID: host.ID(), Addrs: host.Addrs()}
	addrs, err := peer.AddrInfoToP2pAddrs(&ai)
	if err != nil {
		return nil, err
	}
	for _, ma := range addrs {
		fmt.Println(ma)
	}

	return &Node{host: host, pp: pp}, nil
}

// connect connects to a peer with the multiaddress string given and
// returns its peer information.
func (node *Node) connect(ma string) (*peer.AddrInfo, error) {
	addr, err := multiaddr.NewMultiaddr(ma)
	if err != nil {
		return nil, err
	}

	ai, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return nil, err
	}

	if err := node.host.Connect(context.Background(), *ai); err != nil {
		return nil, err
	}

	return ai, nil
}

// Ping connects to a peer using the given multiaddress and sends a ping
// request.
func (node *Node) Ping(peer string) error {
	pi, err := node.connect(peer)
	if err != nil {
		return err
	}

	node.pp.Ping(context.Background(), pi.ID)

	return nil
}
