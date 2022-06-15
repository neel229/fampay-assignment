package db

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	yt "github.com/neel229/fampay-assignment/pkg/youtube"
)

// InsertVideos takes a slice of videos and performs batch insert
func (s *Store) InsertVideos(videos []yt.VideoMetadata) (rowsAffected int, err error) {
	batch := &pgx.Batch{}
	for _, video := range videos {
		uuid, err := uuid.NewUUID()
		if err != nil {
			return 0, ErrUUID
		}
		batch.Queue("insert into videos(id, title, video_id, description, published_at, thumbnail_url) values($1, $2, $3, $4, $5, $6)", uuid, video.Title, video.ID, video.Description, video.PublishedAt, video.ThumbnailURL)
	}
	batchResult := s.pool.SendBatch(context.TODO(), batch)
	for i := 0; i < batch.Len(); i++ {
		ct, err := batchResult.Exec()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		if ct.RowsAffected() != 1 {
			return 0, ErrInsertingVideo
		}
		rowsAffected += int(ct.RowsAffected())
	}
	return
}