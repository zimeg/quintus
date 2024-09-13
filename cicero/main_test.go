package main

import (
	"testing"
	"time"

	ntpclient "github.com/beevik/ntp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimeg/quintus/cicero/pkg/udp"
)

// TestCicero is the complete test for telling the time
//
// It should match main for the most part
func TestCicero(t *testing.T) {
	opts := udp.Options{
		Port:       123,
		PacketSize: 48,
	}
	conn, buff, err := udp.Start(opts)
	require.NoError(t, err)
	go func() {
		addr, err := conn.Read(buff)
		require.NoError(t, err)
		respond(conn, addr, buff)
	}()
	response, err := ntpclient.Time("localhost")
	require.NoError(t, err)
	now := time.Now()
	wait := now.Sub(response)
	assert.Greater(t, wait, 0*time.Millisecond)
	assert.LessOrEqual(t, wait, 1200*time.Millisecond)
}
