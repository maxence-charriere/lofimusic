package main

import (
	"sort"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

var (
	channels = newChannelStore()
)

type channel struct {
	Name        string
	ID          string
	Slug        string
	Description string
}

type channelStore struct {
	channels []channel
}

func newChannelStore() *channelStore {
	rand.Seed(uint64(time.Now().UnixNano()))

	c := []channel{
		{
			Name: "ChilledCow - relax & study",
			ID:   "5qap5aO4i9A",
			Slug: "chilledcow",
		},
		{
			Name: "ChilledCow - sleep & chill",
			ID:   "DWcJFNfaw9c",
			Slug: "chilledcow2",
		},
		{
			Name: "Chillhop Music - jazzy & hip hop",
			ID:   "5yx6BWlEVcY",
			Slug: "chillhopmusic",
		},
		{
			Name: "Chillhop Music - study & relax",
			ID:   "7NOSDKb0HlU",
			Slug: "chillhopmusic2",
		},
		{
			Name: "Tokyo LosT Tracks - サクラチル",
			ID:   "WBfbkPTqUtU",
			Slug: "tokyolosttracks",
		},
		{
			Name: "The Jazz Hop Café",
			ID:   "OVPPOwMpSpQ",
			Slug: "thejazzhopcafe",
		},
	}

	sort.Slice(c, func(a, b int) bool {
		return strings.Compare(c[a].Name, c[b].Name) < 0
	})

	return &channelStore{
		channels: c,
	}
}

func (s *channelStore) Get(slug string) channel {
	for _, c := range s.channels {
		if c.Slug == slug {
			return c
		}
	}

	idx := rand.Intn(len(s.channels))
	return s.channels[idx]
}

func (s *channelStore) Channels() []channel {
	c := make([]channel, len(s.channels))
	copy(c, s.channels)
	return c
}

func (s *channelStore) Slugs() []string {
	slugs := make([]string, len(s.channels))

	for i, c := range s.channels {
		slugs[i] = c.Slug
	}

	return slugs
}
