package track_activities

import "music-app/internal/models/track_activities"

func (r *repository) Create(model track_activities.TrackActivity) error {
	return r.db.Create(&model).Error
}

func (r *repository) Update(model track_activities.TrackActivity) error {
	return r.db.Save(&model).Error
}

func (r *repository) Get(userID uint, spotifyID string) (*track_activities.TrackActivity, error) {
	activity := track_activities.TrackActivity{}
	err := r.db.Where("user_id = ?", userID).Where("spotify_id = ?", spotifyID).First(&activity).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *repository) GetBulk(userID uint, spotifyIDs []string) (map[string]track_activities.TrackActivity, error) {
	activities := make([]track_activities.TrackActivity, 0)
	err := r.db.Where("user_id = ?", userID).Where("spotify_id IN ?", spotifyIDs).Find(&activities).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]track_activities.TrackActivity, 0)
	for _, activity := range activities {
		result[activity.SpotifyID] = activity
	}
	return result, nil
}
