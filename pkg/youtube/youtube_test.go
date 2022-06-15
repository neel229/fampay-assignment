package youtube_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/neel229/fampay-assignment/pkg/utils"
	yt "github.com/neel229/fampay-assignment/pkg/youtube"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

func GetClient() *youtube.Service {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatalf("there was an error loading config file: %v\n", err)
	}
	httpClient := &http.Client{
		Transport: &transport.APIKey{Key: config.GoogleAPIKey},
	}
	client, _ := youtube.New(httpClient)
	return client
}

func TestYTSearch(t *testing.T) {
	client := GetClient()
	sq := yt.SearchQuery{Keyword: "dogs"}
	videos, err := yt.Search(client, sq)
	require.NoError(t, err)
	require.NotEmpty(t, videos)

	for _, video := range videos {
		require.NotEmpty(t, video.Title)
		require.NotEmpty(t, video.ID)
		require.NotEmpty(t, video.PublishedAt)
		require.NotEmpty(t, video.ThumbnailURL)
	}
}
