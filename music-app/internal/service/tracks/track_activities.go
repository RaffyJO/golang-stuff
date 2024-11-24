package tracks

import (
	"fmt"
	"music-app/internal/models/track_activities"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivity(userID uint, request track_activities.TrackActivityRequest) error {
	activity, err := s.trackActivityRepository.Get(userID, request.SpotifyID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error getting track activity from database")
		return err
	}

	if err == gorm.ErrRecordNotFound || activity == nil {
		err = s.trackActivityRepository.Create(track_activities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			IsLiked:   request.IsLiked,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		})
		if err != nil {
			log.Error().Err(err).Msg("Error creating track activity in database")
			return err
		}
		return nil
	}

	activity.IsLiked = request.IsLiked
	err = s.trackActivityRepository.Update(*activity)
	if err != nil {
		log.Error().Err(err).Msg("Error updating track activity in database")
		return err
	}
	return nil
}
