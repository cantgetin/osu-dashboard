package bootstrap

import (
	"net/http"
	"sync"
	"time"
)

type CounterTransport struct {
	Transport http.RoundTripper
	mu        sync.Mutex
	count     int
	success   int
	totalTime time.Duration
}

func (ct *CounterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	ct.mu.Lock()
	defer ct.mu.Unlock()

	if ct.Transport == nil {
		ct.Transport = http.DefaultTransport
	}

	resp, err := ct.Transport.RoundTrip(req)

	ct.count++
	if err == nil && resp != nil && resp.StatusCode < 400 {
		ct.success++
	}
	ct.totalTime += time.Since(start)

	return resp, err
}

func (ct *CounterTransport) RequestCount() int {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	return ct.count
}

func (ct *CounterTransport) SuccessCount() int {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	return ct.success
}

func (ct *CounterTransport) SuccessRate() float64 {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	if ct.count == 0 {
		return 0
	}
	return float64(ct.success) / float64(ct.count) * 100
}

func (ct *CounterTransport) AverageResponseTime() time.Duration {
	ct.mu.Lock()
	defer ct.mu.Unlock()

	if ct.count == 0 {
		return 0
	}
	return ct.totalTime / time.Duration(ct.count)
}

func (ct *CounterTransport) ResetStats() {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	ct.count = 0
	ct.success = 0
	ct.totalTime = 0
}

func NewHTTPClient() http.Client {
	return http.Client{
		Transport: &CounterTransport{},
	}
}
