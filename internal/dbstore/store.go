package dbstore

import (
	"projects/content_service/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBStore struct {
	log  logger.Logger
	db   *pgxpool.Pool
	blog *Blog
}

func New(log logger.Logger, db *pgxpool.Pool) *DBStore {
	return &DBStore{
		log: log,
		db:  db,
	}
}

func (d *DBStore) BlogRepo() *Blog {
	if d.blog == nil {
		d.blog = NewBlog(d.db)
	}
	return d.blog
}
