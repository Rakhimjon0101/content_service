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
func (s *Blog) CreateBlog(ctx context.Context, blog *domains.Blog) error {
	query := "INSERT INTO blogs (id, title, content) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := s.pool.QueryRow(ctx, query, blog.ID, blog.Title, blog.Content).Scan(&blog.ID, &blog.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create blog: %v", err)
	}
	return nil
}

// GetBlogByID get a new blog by it's id.
func (s *Blog) GetBlogByID(ctx context.Context, id string) (domains.Blog, error) {
	blog := domains.Blog{}
	query := "SELECT title, content, created_at FROM blogs WHERE id=$1 AND deleted_at IS NULL"
	err := s.pool.QueryRow(ctx, query, id).Scan(&blog.Title, &blog.Content, &blog.CreatedAt)
	if err != nil {
		return blog, fmt.Errorf("failed to get by id blog: %v", err)
	}
	return blog, nil
}
