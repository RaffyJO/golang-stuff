package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

type SpotifySearchResponse struct {
	Tracks SpotifyTracks `json:"tracks"`
}

type SpotifyRecommendationResponse struct {
	Tracks []SpotifyTrackObject `json:"tracks"`
}

type SpotifyTracks struct {
	Href     string               `json:"href"`
	Limit    int                  `json:"limit"`
	Next     *string              `json:"next"`
	Offset   int                  `json:"offset"`
	Previous *string              `json:"previous"`
	Total    int                  `json:"total"`
	Items    []SpotifyTrackObject `json:"items"`
}

type SpotifyTrackObject struct {
	Album    SpotifyAlbumObject    `json:"album"`
	Artists  []SpotifyArtistObject `json:"artists"`
	Explicit bool                  `json:"explicit"`
	Href     string                `json:"href"`
	ID       string                `json:"id"`
	Name     string                `json:"name"`
}

type SpotifyAlbumObject struct {
	AlbumType   string               `json:"album_type"`
	TotalTracks int                  `json:"total_tracks"`
	Images      []SpotifyAlbumImages `json:"images"`
	Name        string               `json:"name"`
}

type SpotifyAlbumImages struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyArtistObject struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

func (o *Outbound) Search(ctx context.Context, query string, limit, offset int) (*SpotifySearchResponse, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", strconv.Itoa(limit))
	params.Set("offset", strconv.Itoa(offset))

	basePath := "https://api.spotify.com/v1/search"
	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error create request for spotify search")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("Error getting token details for spotify search")
		return nil, err
	}

	bearerToken := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", bearerToken)

	res, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Error executing request for spotify search")
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Error().Msgf("Error getting spotify search. Status code: %d", res.StatusCode)
		return nil, err
	}

	var response SpotifySearchResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling spotify response")
		return nil, err
	}
	return &response, nil
}
