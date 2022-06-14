package youtube

import (
	"log"
	"time"

	youtube "google.golang.org/api/youtube/v3"
)

// Search returns a list of videos based
// on search query params
func Search(client *youtube.Service, sq SearchQuery) ([]VideoMetadata, error) {
	svc := youtube.NewSearchService(client)
	part := []string{"id", "snippet"}
	var call *youtube.SearchListCall
	publishTime := time.Now().Add(time.Duration(-15) * time.Minute).Format(time.RFC3339)
	call = svc.List(part).Q(sq.Keyword).MaxResults(sq.MaxResults).Order("date").Type("video").PublishedAfter(publishTime)
	response, err := call.Do()
	if err != nil {
		log.Printf("there was an error when searching for videos: %v\n", err)
		return []VideoMetadata{}, nil
	}

	videos := make([]VideoMetadata, 0, 10)
	for _, item := range response.Items {
		video := VideoMetadata{
			Title:        item.Snippet.Title,
			ID:           item.Id.VideoId,
			Description:  item.Snippet.Description,
			PublishedAt:  item.Snippet.PublishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
		}
		videos = append(videos, video)
	}
	return videos, nil
}
