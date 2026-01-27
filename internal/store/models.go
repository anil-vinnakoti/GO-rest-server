package store

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Source    url.URL   `json:"source"`
	Tags      []string  `json:"tags"`
}
