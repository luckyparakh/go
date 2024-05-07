package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}

// Transport is an interface that handles the communication between nodes in the network.
// It can be form like UDP, TCP, Websocket etc.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
