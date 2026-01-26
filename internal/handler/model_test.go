package handler_test

import (
	"testing"

	"github.com/anil-vinnakoti/newsapi/internal/handler"
)

func TestNewsPostRequestBody_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		req      handler.NewsPostRequestBody
		expected bool
	}{
		{
			name:     "author empty",
			req:      handler.NewsPostRequestBody{},
			expected: true,
		},
		{
			name: "title empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
			},
			expected: true,
		},
		{
			name: "summary empty",
			req: handler.NewsPostRequestBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expected: true,
		},
		{
			name: "time empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid time",
			},
			expected: true,
		},
		{
			name: "source empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
			},
			expected: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://test-news.com",
			},
			expected: true,
		},
		{
			name: "validate",
			req: handler.NewsPostRequestBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2024-04-07T05:13:27+00:00",
				Source:    "https://test-news.com",
				Tags:      []string{"test-tag"},
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.req.Validate()
			if testCase.expected && err == nil {
				t.Fatalf("expected error but got nil")
			}
			if !testCase.expected && err != nil {
				t.Fatalf("expected nil but got error: %v", err)
			}
		})
	}
}
