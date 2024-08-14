package p2p

// interface that represents the remote Node.
type Peer interface {
}

// Anything that can handle communication between the Nodes in the
// network(TCP, UDP, websockets...)
type Transport interface {
	ListenAndAccept() error
}
