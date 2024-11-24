package spotify

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *Outbound) GetTokenDetails() (string, string, error) {
	if o.AccessToken == "" || o.ExpiredAt.Before(time.Now()) {
		err := o.generateToken()
		if err != nil {
			return "", "", err
		}
	}

	return o.AccessToken, o.TokenType, nil
}

func (o *Outbound) generateToken() error {
	formData := url.Values{}
	formData.Add("grant_type", "client_credentials")
	formData.Add("client_id", o.cfg.Spotify.ClientID)
	formData.Add("client_secret", o.cfg.Spotify.ClientSecret)

	encodedUrl := formData.Encode()

	req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(encodedUrl))
	if err != nil {
		log.Error().Err(err).Msg("Error creating request for spotify token")
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Error getting spotify token")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Error().Msgf("Error getting spotify token. Status code: %d", res.StatusCode)
		return err
	}

	var response SpotifyTokenResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling spotify response")
		return err
	}

	o.AccessToken = response.AccessToken
	o.TokenType = response.TokenType
	o.ExpiredAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)

	return nil
}
