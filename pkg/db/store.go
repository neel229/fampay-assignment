package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewStore creates a new store instance
func NewStore(url string) (store *Store, err error) {
	pool, err := pgxpool.Connect(context.Background(), url)
	if err := pool.Ping(context.TODO()); err != nil {
		log.Println(err)
	}
	store = &Store{
		pool: pool,
	}
	return
}

func (s *Store) CloseConn() {
	s.pool.Close()
}
