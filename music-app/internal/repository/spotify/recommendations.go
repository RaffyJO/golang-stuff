package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (o *Outbound) GetRecommendations(limit int, trackID string) (*SpotifyRecommendationResponse, error) {
	params := url.Values{}
	params.Set("limit", strconv.Itoa(limit))
	params.Set("market", "ID")
	params.Set("seed_tracks", trackID)

	basePath := "https://api.spotify.com/v1/recommendations"
	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error create request for spotify recommendations")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("Error getting token details for spotify recommendations")
		return nil, err
	}

	bearerToken := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", bearerToken)

	res, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Error executing request for spotify recommendations")
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Error().Msgf("Error getting spotify recommendations. Status code: %d", res.StatusCode)
		return nil, err
	}

	var response SpotifyRecommendationResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling spotify response")
		return nil, err
	}
	return &response, nil
}
