package pwa

import (
	"testing"

	"github.com/maxence-charriere/lofimusic/pkg/radio"
	"github.com/stretchr/testify/require"
)

func TestVideoStoreLeastWatchedVideos(t *testing.T) {
	store := newVideoStore()
	expectedLen := len(radio.List())

	lwv := store.leastWatchedVideos()
	require.Len(t, lwv, expectedLen)

	store.inc(store.videoBySlug("lofigirl"))
	lwv = store.leastWatchedVideos()
	expectedLen--
	require.Len(t, lwv, expectedLen)
}
