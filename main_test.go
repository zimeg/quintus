package main

import (
	"testing"
	"time"

	ntpclient "github.com/beevik/ntp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimeg/quintus/pkg/now"
	"github.com/zimeg/quintus/pkg/udp"
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
	response, err := ntpclient.Time("localhost:12321")
	require.NoError(t, err)
	wait := now.Moment(time.Now().UTC()).Epoch() - uint64(response.Unix())
	assert.GreaterOrEqual(t, wait, uint64(0))
	assert.LessOrEqual(t, wait, uint64(1200))
}
