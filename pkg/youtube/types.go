package youtube

// SearchQuery holds the different
// search parameters
type SearchQuery struct {
	Keyword    string `json:"keyword"`
	MaxResults int64  `json:"maxResults"`
}

// VideoMetadata contains data of a video
type VideoMetadata struct {
	Title        string `json:"title"`
	ID           string `json:"videoId"`
	Description  string `json:"description"`
	PublishedAt  string `json:"publishedAt"`
	ThumbnailURL string `json:"thumbnailURL"`
}
