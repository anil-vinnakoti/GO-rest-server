package handler

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/anil-vinnakoti/newsapi/internal/news"
	"github.com/google/uuid"
)

type NewsPostRequestBody struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"created_at"`
	Content   string    `json:"content"`
	Source    string    `json:"source"`
	Tags      []string  `json:"tags"`
}

func (n NewsPostRequestBody) Validate() (record *news.Record, errs error) {
	if n.Author == "" {
		errs = errors.Join(errs, fmt.Errorf("author is empty: %s", n.Author))
	}
	if n.Title == "" {
		errs = errors.Join(errs, fmt.Errorf("title is empty: %s", n.Title))
	}
	if n.Summary == "" {
		errs = errors.Join(errs, fmt.Errorf("summary is empty: %s", n.Summary))
	}
	if n.Content == "" {
		errs = errors.Join(errs, fmt.Errorf("content is empty"))
	}

	t, err := time.Parse(time.RFC3339, n.CreatedAt)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	if n.Source == "" {
		errs = errors.Join(errs, fmt.Errorf("source is empty: %s", n.Source))
	}
	url, err := url.Parse(n.Source)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	if len(n.Tags) == 0 {
		errs = errors.Join(errs, errors.New("tags cannot be empty"))
	}

	if errs != nil {
		return record, errs
	}
	return &news.Record{
		ID:        n.ID,
		Author:    n.Author,
		Title:     n.Author,
		Content:   n.Content,
		Summary:   n.Summary,
		CreatedAt: t,
		Source:    url.String(),
		Tags:      n.Tags,
	}, nil
}
