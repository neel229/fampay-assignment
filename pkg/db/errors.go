package db

import "errors"

var (
	ErrUUID           = errors.New("error generating uuid")
	ErrInsertingVideo = errors.New("error inserting video")
)
