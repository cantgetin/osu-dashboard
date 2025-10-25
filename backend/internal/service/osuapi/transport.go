package osuapi

import (
	"time"
)

type TransportStats struct {
	RequestCount    int
	AvgResponseTime time.Duration
	SuccessRate     float64
}

func (s *Service) GetTransportStats() TransportStats {
	return TransportStats{
		RequestCount:    s.GetOutgoingRequestCount(),
		AvgResponseTime: s.AverageResponseTime(),
		SuccessRate:     s.SuccessRate(),
	}
}

func (s *Service) GetOutgoingRequestCount() int {
	return s.httpClient.RequestCount()
}

func (s *Service) AverageResponseTime() time.Duration {
	return s.httpClient.AverageResponseTime()
}

func (s *Service) SuccessRate() float64 {
	return s.httpClient.SuccessRate()
}

func (s *Service) ResetStats() {
	s.httpClient.ResetStats()
}
