package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	ListenAddr := ":4000"
	tr := NewTCPTransport(ListenAddr)
	assert.Equal(t, tr.ListenAddr, ListenAddr)
	assert.Nil(t, tr.ListenAndAccept())
	select {}
}
