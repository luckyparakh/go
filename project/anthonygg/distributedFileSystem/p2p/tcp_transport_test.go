package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	options := TCPTransportOpts{
		ListenAddress: ":3000",
		Decoder: DefaultDecode{},
		HandShaker: NOPHandShake,
	}
	nt := NewTCPTransport(options)
	assert.Equal(t, nt.TCPTransportOpts.ListenAddress, options.ListenAddress)

	assert.Nil(t, nt.ListenAndAccept())
}
