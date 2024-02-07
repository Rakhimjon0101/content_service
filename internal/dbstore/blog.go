package dbstore

import (
	"context"
	"fmt"
	"projects/content_service/internal/domains"

	"github.com/jackc/pgx/v5/pgxpool"
)

// BlogStorage handles CRUD operations for blogs.
type Blog struct {
	pool *pgxpool.Pool
}

// NewBlogStorage creates a new BlogStorage instance.
func NewBlog(pool *pgxpool.Pool) *Blog {
	return &Blog{pool: pool}
}

// CreateBlog creates a new blog post.
func (s *Blog) CreateBlog(blog *domains.Blog) error {
	query := "INSERT INTO blogs (title, content) VALUES ($1, $2) RETURNING id, created_at"
	err := s.pool.QueryRow(context.Background(), query, blog.Title, blog.Content).Scan(&blog.ID, &blog.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create blog: %v", err)
	}
	return nil
}
