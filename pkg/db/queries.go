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

func (s *Store) GetVideos(limit, offset int) ([]yt.VideoMetadata, error) {
	query := "SELECT * FROM videos ORDER BY published_at DESC LIMIT $1 OFFSET $2"
	rows, err := s.pool.Query(context.TODO(), query, limit, offset)
	if err != nil {
		log.Println(err)
		return []yt.VideoMetadata{}, err
	}
	defer rows.Close()
	var videos []yt.VideoMetadata
	for rows.Next() {
		var video yt.VideoMetadata
		if err := rows.Scan(nil, &video.Title, &video.ID, &video.Description, &video.PublishedAt, &video.ThumbnailURL); err != nil {
			log.Println(err)
			return []yt.VideoMetadata{}, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (s *Store) SearchWithTitle(title string) (yt.VideoMetadata, error) {
	query := "SELECT * FROM videos WHERE title = $1"
	row := s.pool.QueryRow(context.TODO(), query, title)
	var video yt.VideoMetadata
	if err := row.Scan(nil, &video.Title, &video.ID, &video.Description, &video.PublishedAt, &video.ThumbnailURL); err != nil {
		log.Println(err)
		return yt.VideoMetadata{}, err
	}
	return video, nil
}

func (s *Store) SearchWithDescription(title string) (yt.VideoMetadata, error) {
	query := "SELECT * FROM videos WHERE description = $1"
	row := s.pool.QueryRow(context.TODO(), query, title)
	var video yt.VideoMetadata
	if err := row.Scan(nil, &video.Title, &video.ID, &video.Description, &video.PublishedAt, &video.ThumbnailURL); err != nil {
		log.Println(err)
		return yt.VideoMetadata{}, err
	}
	return video, nil
}
