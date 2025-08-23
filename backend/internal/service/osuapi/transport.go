package osuapi

import (
	"playcount-monitor-backend/internal/bootstrap"
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
	return s.httpClient.Transport.(*bootstrap.CounterTransport).RequestCount()
}

func (s *Service) AverageResponseTime() time.Duration {
	return s.httpClient.Transport.(*bootstrap.CounterTransport).AverageResponseTime()
}

func (s *Service) SuccessRate() float64 {
	return s.httpClient.Transport.(*bootstrap.CounterTransport).SuccessRate()
}

func (s *Service) ResetStats() {
	s.httpClient.Transport.(*bootstrap.CounterTransport).ResetStats()
}
