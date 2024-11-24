package tracks

import (
	"context"
	"music-app/internal/models/spotify"
	"music-app/internal/models/track_activities"
	spotifyRepo "music-app/internal/repository/spotify"

	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("Error getting spotify search")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for i, item := range trackDetails.Tracks.Items {
		trackIDs[i] = item.ID
	}

	trackActivities, err := s.trackActivityRepository.GetBulk(userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("Error getting track activities")
		return nil, err
	}

	return modelToResponse(trackDetails, trackActivities), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]track_activities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
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

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Total:  data.Tracks.Total,
		Items:  items,
	}
}
