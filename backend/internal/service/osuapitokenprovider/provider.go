package osuapitokenprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (p *TokenProvider) GetToken(ctx context.Context) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.validUntil.Add(-time.Minute).After(time.Now().UTC()) {
		return p.token, nil
	}

	// curl --request POST \
	//    "https://osu.ppy.sh/oauth/token" \
	//    --header "Accept: application/json" \
	//    --header "Content-Type: application/x-www-form-urlencoded" \
	//    --data "client_id=1&client_secret=clientsecret&grant_type=client_credentials&scope=public"

	values := url.Values{}
	values.Set("client_id", p.cfg.OsuAPIClientID)
	values.Set("client_secret", p.cfg.OsuAPIClientSecret)
	values.Set("grant_type", "client_credentials")
	values.Set("scope", "public")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.cfg.OsuOAuthHost, strings.NewReader(values.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to invoke request: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Token   string        `json:"access_token"`
		Expires time.Duration `json:"expires_in"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	p.token = data.Token
	p.validUntil = time.Now().UTC().Add(time.Second * data.Expires)

	return p.token, nil
}
