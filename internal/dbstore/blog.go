package dbstore

import (
	"context"
	"fmt"

	"projects/content_service/internal/domains"
	"projects/content_service/internal/errs"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Blog handles CRUD operations for blogs.
type Blog struct {
	pool *pgxpool.Pool
}

// NewBlog creates a new BlogStorage instance.
func NewBlog(pool *pgxpool.Pool) *Blog {
	return &Blog{pool: pool}
}

// CreateBlog creates a new blog post.
func (s *Blog) CreateBlog(ctx context.Context, blog *domains.Blog) error {
	query := "INSERT INTO blogs (id, title, content) VALUES ($1, $2, $3) RETURNING id, created_at"
	err := s.pool.QueryRow(ctx, query, blog.ID, blog.Title, blog.Content).Scan(&blog.ID, &blog.CreatedAt)
	if err != nil {
		return fmt.Errorf("blogRepo failed to create blog: %v", err)
	}
	return nil
}

// GetBlogByID get a new blog by it's id.
func (s *Blog) GetBlogByID(ctx context.Context, id string) (domains.Blog, error) {
	blog := domains.Blog{}
	query := "SELECT title, content, created_at FROM blogs WHERE id=$1 AND deleted_at IS NULL"
	err := s.pool.QueryRow(ctx, query, id).Scan(&blog.Title, &blog.Content, &blog.CreatedAt)
	if err != nil {
		return blog, fmt.Errorf("blogRepo failed to get by id blog: %v", err)
	}
	return blog, nil
}

// GetBlogs get all blogs
func (s *Blog) GetBlogs(ctx context.Context) ([]domains.Blog, error) {
	var (
		blogs []domains.Blog
		query = "SELECT id, title, content, created_at, updated_at FROM blogs WHERE deleted_at IS NULL"
	)
	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("blogRepo failed to get by id blog: %v", err)
	}

	for rows.Next() {
		var blog domains.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, fmt.Errorf("blogRepo.Get: Scan --%s %w", err.Error(), errs.ErrNotFound)
		}
		blogs = append(blogs, blog)
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("blogRepo.Get: rows.Err-- %w", err)
	}
	return blogs, nil
}
