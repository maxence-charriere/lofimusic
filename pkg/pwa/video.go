package pwa

import (
	"sort"
	"sync"

	"github.com/maxence-charriere/lofimusic/pkg/radio"
)

type videoCount struct {
	slug  string
	count int
}

type videoStore struct {
	mu       sync.Mutex
	videos   map[string]radio.Video
	counters []videoCount
}

func newVideoStore() *videoStore {
	rl := radio.List()
	videos := make(map[string]radio.Video, len(rl))
	counters := make([]videoCount, len(rl))

	for i, v := range rl {
		videos[v.Slug] = v
		counters[i].slug = v.Slug
	}

	return &videoStore{
		videos:   videos,
		counters: counters,
	}
}

func (s *videoStore) leastWatchedVideos() []radio.Video {
	s.mu.Lock()
	defer s.mu.Unlock()

	sort.Slice(s.counters, func(a, b int) bool {
		return s.counters[a].count < s.counters[b].count
	})

	lowestCount := s.counters[0].count

	var videos []radio.Video
	for i := 0; i < len(s.counters) && s.counters[i].count == lowestCount; i++ {
		videos = append(videos, s.videoBySlug(s.counters[i].slug))
	}
	return videos
}

func (s *videoStore) videoBySlug(slug string) radio.Video {
	return s.videos[slug]
}

func (s *videoStore) inc(v radio.Video) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.counters {
		if s.counters[i].slug == v.Slug {
			s.counters[i].count++
			return
		}
	}
}
