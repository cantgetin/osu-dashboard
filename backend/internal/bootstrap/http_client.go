package bootstrap

import (
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
	"sync"
	"time"
)

func NewHTTPClient() *CustomHTTPClient {
	c := retryablehttp.NewClient()
	c.RetryMax = 3
	c.RetryWaitMin = 3 * time.Second
	c.RetryWaitMax = 10 * time.Second

	return NewCustomClient(c.StandardClient())
}

// TODO: move somewhere else
type CustomHTTPClient struct {
	*http.Client
	*Stats

	mu *sync.Mutex
}

func NewCustomClient(client *http.Client) *CustomHTTPClient {
	return &CustomHTTPClient{
		Client: client,
		mu:     &sync.Mutex{},
		Stats:  &Stats{},
	}
}

func (cc *CustomHTTPClient) Do(req *http.Request) (*http.Response, error) {
	start := time.Now()
	cc.mu.Lock()
	defer cc.mu.Unlock()

	resp, err := cc.Do(req)

	cc.IncreaseCount()
	if err == nil && resp != nil && resp.StatusCode < http.StatusBadRequest {
		cc.IncreaseSuccessCount()
	}
	cc.IncreaseTotalTime(time.Since(start))
	return resp, err
}
