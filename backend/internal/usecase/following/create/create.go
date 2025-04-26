package followingcreate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"playcount-monitor-backend/internal/database/repository/model"
	"playcount-monitor-backend/internal/database/txmanager"
	"strings"
	"time"
)

func (uc *UseCase) Create(
	ctx context.Context,
	code string,
) error {
	token, err := uc.exchangeCodeForToken(code)
	if err != nil {
		return fmt.Errorf("failed to exchange code for access token: %w", err)
	}

	user, err := getUserInfo(token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to get user info: %w", err)
	}

	following := &model.Following{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: time.Now().UTC(),
	}
	if following.ID <= 0 || following.Username == "" {
		return fmt.Errorf("user id or username is empty: %w", err)
	}

	txErr := uc.txm.ReadWrite(ctx, func(ctx context.Context, tx txmanager.Tx) error {
		f, err := uc.following.Get(ctx, tx, following.ID)
		if f != nil {
			return nil // assume it's already added
		}

		err = uc.following.Create(ctx, tx, following)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		return txErr
	}

	go func() {
		// start tracking background process
		trackCtx := context.Background()
		err = uc.track.TrackSingleFollowing(trackCtx, following)
		if err != nil {
			uc.lg.Printf("failed to track specific user %v", err)
		}
	}()

	return nil
}

// TODO move below to osu api service
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (uc *UseCase) exchangeCodeForToken(code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", uc.cfg.OsuAPIClientID)
	data.Set("client_secret", uc.cfg.OsuAPIClientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", uc.cfg.OsuAPIRedirectURI)

	req, err := http.NewRequest("POST", "https://osu.ppy.sh/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("oauth server returned %d: %s", resp.StatusCode, body)
	}

	var token TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func getUserInfo(accessToken string) (*UserResponse, error) {
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/v2/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("api server returned %d: %s", resp.StatusCode, body)
	}

	var user UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
