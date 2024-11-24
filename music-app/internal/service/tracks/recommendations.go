package tracks

import (
	"music-app/internal/models/spotify"
	"music-app/internal/models/track_activities"
	spotifyRepo "music-app/internal/repository/spotify"

	"github.com/rs/zerolog/log"
)

func (s *service) GetRecommendations(userID uint, limit int, trackID string) (*spotify.RecommendationResponse, error) {
	trackDetails, err := s.spotifyOutbound.GetRecommendations(limit, trackID)
	if err != nil {
		log.Error().Err(err).Msg("Error getting spotify recommendations")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks))
	for i, item := range trackDetails.Tracks {
		trackIDs[i] = item.ID
	}

	trackActivities, err := s.trackActivityRepository.GetBulk(userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("Error getting track activities")
		return nil, err
	}

	return modelToRecommendationResponse(trackDetails, trackActivities), nil
}

func modelToRecommendationResponse(data *spotifyRepo.SpotifyRecommendationResponse, mapTrackActivities map[string]track_activities.TrackActivity) *spotify.RecommendationResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks {
		artistsName := make([]string, len(item.Artists))
		for i, artist := range item.Artists {
			artistsName[i] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for i, image := range item.Album.Images {
			imageUrls[i] = image.Url
		}

		items = append(items, spotify.SpotifyTrackObject{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesUrl:   imageUrls,
			AlbumName:        item.Album.Name,

			ArtistsName: artistsName,

			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,
			IsLiked:  mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.RecommendationResponse{
		Items: items,
	}
}
