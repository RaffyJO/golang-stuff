package tracks

import (
	"context"
	"music-app/internal/models/track_activities"
	"music-app/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type trackActivityRepository interface {
	Create(model track_activities.TrackActivity) error
	Update(model track_activities.TrackActivity) error
	Get(userID uint, spotifyID string) (*track_activities.TrackActivity, error)
	GetBulk(userID uint, spotifyID []string) (map[string]track_activities.TrackActivity, error)
}

type service struct {
	spotifyOutbound         spotifyOutbound
	trackActivityRepository trackActivityRepository
}

func NewService(spotifyOutbound spotifyOutbound, trackActivityRepository trackActivityRepository) *service {
	return &service{
		spotifyOutbound:         spotifyOutbound,
		trackActivityRepository: trackActivityRepository,
	}
}
