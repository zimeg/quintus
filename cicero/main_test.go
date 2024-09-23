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
		Port:       12321,
		PacketSize: 48,
	}
	conn, buff, err := udp.Start(opts)
	require.NoError(t, err)
	go func() {
		addr, err := conn.Read(buff)
		require.NoError(t, err)
		respond(conn, addr, buff)
	}()
	now := time.Now()
	response, err := ntpclient.Time("localhost:12321")
	require.NoError(t, err)
	wait := now.Sub(response)
	assert.Greater(t, int64(0), wait.Milliseconds())
	assert.LessOrEqual(t, wait.Milliseconds()%10000, int64(1200))
}
