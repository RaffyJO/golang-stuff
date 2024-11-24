package spotify

import (
	"music-app/internal/configs"
	"music-app/pkg/httpclient"
	"time"
)

type Outbound struct {
	cfg         *configs.Config
	client      httpclient.HTTPClient
	AccessToken string
	TokenType   string
	ExpiredAt   time.Time
}

func NewSpotifyOutbound(cfg *configs.Config, client httpclient.HTTPClient) *Outbound {
	return &Outbound{
		cfg:    cfg,
		client: client,
	}
}
