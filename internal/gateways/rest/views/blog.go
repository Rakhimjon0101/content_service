package views

import "time"

type Blog struct {
	ID        string     `json:"id,omitempty"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAT *time.Time `json:"updated_at,omitempty"`
	DeletedAT *time.Time `json:"deleted_at,omitempty"`
}
