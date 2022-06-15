package youtube_test

import (
	"log"
	"testing"

	"github.com/neel229/fampay-assignment/pkg/utils"
	yt "github.com/neel229/fampay-assignment/pkg/youtube"
	"github.com/stretchr/testify/require"
)

func TestYTSearch(t *testing.T) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatalf("there was an error loading config file: %v\n", err)
	}
	videos, err := yt.Search(config.GoogleAPIKey, "cats")
	require.NoError(t, err)
	require.NotEmpty(t, videos)

	for _, video := range videos {
		require.NotEmpty(t, video.Title)
		require.NotEmpty(t, video.ID)
		require.NotEmpty(t, video.PublishedAt)
		require.NotEmpty(t, video.ThumbnailURL)
	}
}
