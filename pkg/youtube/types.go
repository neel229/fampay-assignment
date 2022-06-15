package youtube

import "time"

// SearchQuery holds the different
// search parameters
type SearchQuery struct {
	Keyword string `json:"keyword"`
}

// VideoMetadata contains data of a video
type VideoMetadata struct {
	Title        string    `json:"title"`
	ID           string    `json:"videoId"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `json:"publishedAt"`
	ThumbnailURL string    `json:"thumbnailURL"`
}
