package bootstrap

import (
	"sync"
	"time"
)

type Stats struct {
	count     int
	success   int
	mu        sync.Mutex
	totalTime time.Duration
}

func (s *Stats) IncreaseCount() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *Stats) IncreaseSuccessCount() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.success++
}

func (s *Stats) IncreaseTotalTime(t time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalTime += t
}

func (s *Stats) RequestCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func (s *Stats) SuccessCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.success
}

func (s *Stats) SuccessRate() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.count == 0 {
		return 0
	}
	return float64(s.success) / float64(s.count) * 100
}

func (s *Stats) AverageResponseTime() time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.count == 0 {
		return 0
	}
	return s.totalTime / time.Duration(s.count)
}

func (s *Stats) ResetStats() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = 0
	s.success = 0
	s.totalTime = 0
}
