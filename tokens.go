package slack

import (
	"context"
	"net/url"
)

// RotateTokens exchanges a refresh token for a new app configuration token
func (api *Client) RotateTokens(refreshToken string) (*TokenResponse, error) {
	return api.RotateTokensContext(context.Background(), refreshToken)
}

// RotateTokensContext exchanges a refresh token for a new app configuration token with a custom context
func (api *Client) RotateTokensContext(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	if refreshToken == "" {
		refreshToken = api.configRefreshToken
	}

	values := url.Values{
		"refresh_token": {refreshToken},
	}

	response := &TokenResponse{}
	err := api.getMethod(ctx, "tooling.tokens.rotate", api.configToken, values, response)
	if err != nil {
		return nil, err
	}

	return response, response.Err()
}

type TokenResponse struct {
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TeamId       string `json:"team_id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	IssuedAt     uint64 `json:"iat,omitempty"`
	ExpiresAt    uint64 `json:"exp,omitempty"`
	SlackResponse
}
