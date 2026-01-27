package handler_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/anil-vinnakoti/newsapi/internal/handler"
	"github.com/anil-vinnakoti/newsapi/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewsPostRequestBody_Validate(t *testing.T) {
	type expectations struct {
		err  string
		news store.News
	}
	testCases := []struct {
		name string
		req  handler.NewsPostRequestBody
		expectations
	}{
		{
			name:         "author empty",
			req:          handler.NewsPostRequestBody{},
			expectations: expectations{err: "author is empty"},
		},
		{
			name: "title empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
			},
			expectations: expectations{err: "title is empty"},
		},
		{
			name: "content empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectations: expectations{err: "content is empty"},
		},
		{
			name: "summary empty",
			req: handler.NewsPostRequestBody{
				Author:  "test-author",
				Title:   "test-title",
				Content: "test-content",
			},
			expectations: expectations{err: "summary is empty"},
		},
		{
			name: "time empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				Content:   "test-content",
				CreatedAt: "invalid time",
			},
			expectations: expectations{err: `parsing time "invalid time"`},
		},
		{
			name: "source empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				Content:   "test-content",
				CreatedAt: "2024-04-07T05:13:27+00:00",
			},
			expectations: expectations{err: "source is empty"},
		},
		{
			name: "tags empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Content:   "test-content",
				Source:    "https://test-news.com",
			},
			expectations: expectations{err: "tags cannot be empty"},
		},
		{
			name: "validate",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://test-news.com",
				Content:   "test-content",
				Tags:      []string{"test-tag"},
			},
			expectations: expectations{err: "", news: store.News{
				Author:  "test-author",
				Title:   "test-title",
				Summary: "test-summary",
				Content: "test-content",
				Tags:    []string{"test-tag"}}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := testCase.req.Validate()

			if testCase.expectations.err != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), testCase.expectations.err)
			} else {
				assert.NoError(t, err)

				parsedTime, parseTimeErr := time.Parse(time.RFC3339, testCase.req.CreatedAt)
				require.NoError(t, parseTimeErr)
				testCase.expectations.news.CreatedAt = parsedTime

				parsedSource, parseSourceErr := url.Parse(testCase.req.Source)
				require.NoError(t, parseSourceErr)
				testCase.expectations.news.Source = *parsedSource

			}
		})
	}
}
