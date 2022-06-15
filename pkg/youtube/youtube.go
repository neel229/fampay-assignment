package youtube

import (
	"log"
	"net/http"
	"time"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

// Search returns a list of videos based
// on search query params
func Search(apiKey, keyword string) ([]VideoMetadata, error) {
	httpClient := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	client, err := youtube.New(httpClient)
	if err != nil {
		log.Printf("error creating yt client: %v\n", err)
		return []VideoMetadata{}, err
	}
	svc := youtube.NewSearchService(client)
	part := []string{"id", "snippet"}
	publishTime := time.Now().Add(time.Duration(-20) * time.Minute).Format(time.RFC3339)
	var call *youtube.SearchListCall
	call = svc.List(part).Q(keyword).Order("date").Type("video").PublishedAfter(publishTime)
	response, err := call.Do()
	if err != nil {
		log.Printf("there was an error when searching for videos: %v\n", err)
		return []VideoMetadata{}, ErrSearchingVideos
	}

	videos := make([]VideoMetadata, 0, 10)
	for _, item := range response.Items {
		time, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			log.Println(err)
		}
		video := VideoMetadata{
			Title:        item.Snippet.Title,
			ID:           item.Id.VideoId,
			Description:  item.Snippet.Description,
			PublishedAt:  time,
			ThumbnailURL: item.Snippet.Thumbnails.Default.Url,
		}
		videos = append(videos, video)
	}
	return videos, nil
}
