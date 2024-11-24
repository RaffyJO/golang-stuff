package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Total  int                  `json:"total"`
	Items  []SpotifyTrackObject `json:"items"`
}

type SpotifyTrackObject struct {
	AlbumType        string   `json:"album_type"`
	AlbumTotalTracks int      `json:"album_total_tracks"`
	AlbumImagesUrl   []string `json:"album_images_url"`
	AlbumName        string   `json:"album_name"`

	ArtistsName []string `json:"artists_name"`

	Explicit bool   `json:"explicit"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsLiked  *bool  `json:"is_liked"`
}

type RecommendationResponse struct {
	Items []SpotifyTrackObject `json:"items"`
}
