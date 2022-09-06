// ping defines a simple liveness check protocol where the node will
// send a UTF-8 encoded string "PING" and the peer responds with "PONG"
// in the same stream.
package ping

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const (
	id   = "/juno/ping/0.1.0"
	name = "gojuno.xyz.ping"
)

const (
	// pingLen is the size of a single ping message in bytes.
	pingLen = 4
	timeOut = 30 * time.Second
)

// Protocol represents a ping protocol.
type Protocol struct {
	host host.Host
}

// Register registers the ping protocol with the given host.
func Register(h host.Host) *Protocol {
	p := &Protocol{host: h}
	h.SetStreamHandler(id, p.Pong)
	return p
}

// Pong contains the logic of what happens when a node receives a ping
// request. This is the function that gets registered with the host as a
// network stream handler.
func (p *Protocol) Pong(s network.Stream) {
	if err := p.pong(s); err != nil {
		// XXX: Log error at debug?
		s.Reset()
	}
	s.Close()
}

func (p *Protocol) pong(s network.Stream) error {
	// Attach stream to ping service.
	if err := s.Scope().SetService(name); err != nil {
		return err
	}

	// Reserve memory for ping stream.
	if err := s.Scope().ReserveMemory(pingLen, network.ReservationPriorityAlways); err != nil {
		return err
	}
	defer s.Scope().ReleaseMemory(pingLen)

	// NOTE: There is a ping timeout and ping duration. See the following
	// PR for details https://github.com/libp2p/go-libp2p/pull/1358.
	s.SetReadDeadline(time.Now().Add(timeOut))

	// XXX: Use sync.Pool?
	msg := make([]byte, pingLen)
	_, err := io.ReadFull(s, msg)
	if err != nil {
		return err
	}

	// DEBUG.
	fmt.Printf("Read: %s.\n", msg)

	// Write back to stream.
	_, err = s.Write([]byte("PONG"))

	return err
}

// Ping sends a ping request to the peer with peerId pi.
func (p *Protocol) Ping(ctx context.Context, peer peer.ID) error {
	s, err := p.host.NewStream(context.Background(), peer, id)
	if err != nil {
		return err
	}

	// Write to stream.
	if _, err := s.Write([]byte("PING")); err != nil {
		return err
	}

	// Read in the response from the stream after sending the message
	// above.
	msg := make([]byte, pingLen)
	_, err = io.ReadFull(s, msg)
	if err != nil {
		return err
	}

	// DEBUG.
	fmt.Printf("Read: %s.\n", msg)

	return nil
}
