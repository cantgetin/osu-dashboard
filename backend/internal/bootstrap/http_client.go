package bootstrap

import (
	"net/http"
	"sync"
)

type CounterTransport struct {
	Transport http.RoundTripper
	mu        sync.Mutex
	count     int
}

func (ct *CounterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	ct.count++

	if ct.Transport == nil {
		ct.Transport = http.DefaultTransport
	}

	return ct.Transport.RoundTrip(req)
}

func (ct *CounterTransport) RequestCount() int {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	return ct.count
}

func (ct *CounterTransport) ResetCount() {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	ct.count = 0
}

func NewHTTPClient() http.Client {
	return http.Client{
		Transport: &CounterTransport{},
	}
}
