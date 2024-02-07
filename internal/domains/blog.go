package domains

import (
	"time"
)

// Blog represents a blog post.
type Blog struct {
	ID        string     `db:"id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAT *time.Time `db:"updated_at"`
	DeletedAT *time.Time `db:"deleted_at"`
}
