package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}
