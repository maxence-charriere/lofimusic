package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerStatus(t *testing.T) {
	var p playerState

	err := p.load()
	require.NoError(t, err)
	require.Equal(t, 100, p.LastNonZeroVolume)
	require.Equal(t, 100, p.Volume)
	require.False(t, p.Muted)

	p.Volume = 50
	p.Muted = true
	err = p.save()
	require.NoError(t, err)

	p = playerState{}
	err = p.load()
	require.NoError(t, err)
	require.Equal(t, 100, p.LastNonZeroVolume)
	require.Equal(t, 50, p.Volume)
	require.True(t, p.Muted)
}
