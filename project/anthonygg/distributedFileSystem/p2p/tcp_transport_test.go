package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddress := ":4000"
	nt := NewTCPTransport(listenAddress)
	assert.Equal(t, nt.listenAddress, listenAddress)

	assert.Nil(t, nt.ListenAndAccept())
}
