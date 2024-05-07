package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over established connection.
type TCPPeer struct {
	// conn is underlying connection of peer
	conn net.Conn

	// if we dial and retrieve the connection => outbound==true
	// if we accept and receive the connection => outbound==true
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// Close implements the Peer interface
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	HandShaker    HandShakeFunc
	Decoder       Decoder
	ListenAddress string
	OnPeer        func(Peer) error
}
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcCh    chan RPC

	mu sync.RWMutex
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcCh:            make(chan RPC),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

// Consume implements Transport interface, which will return read only channel for reading incoming message
// coming from Peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcCh
}
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	defer func() {
		fmt.Printf("Dropping Peer %s\n", err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn, true)
	if err = t.HandShaker(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP hand shake  error: %s\n", err)
		return
	}
	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}
	rpc := RPC{}
	for {
		if err = t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			return
		}
		rpc.From = conn.RemoteAddr()
		t.rpcCh <- rpc
	}
}
