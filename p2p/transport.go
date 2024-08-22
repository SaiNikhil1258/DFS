package p2p

import "net"

// interface that represents the remote Node.
type Peer interface {
	net.Conn
	Send([]byte) error
}

// Anything that can handle communication between the Nodes in the
// network(TCP, UDP, websockets...)
type Transport interface {
	ListenAndAccept() error
	Dial(string) error
	Consume() <-chan RPC
	Close() error
	// ListenAddr() string
}
