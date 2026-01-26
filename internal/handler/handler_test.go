package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anil-vinnakoti/newsapi/internal/handler"
	"github.com/google/uuid"
)

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "db error",
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusInternalServerError,
		}, {
			name:           "success",
			store:          mockNewsStore{},
			expectedStatus: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			// Act
			handler.GetAllNews(testCase.store)(w, r)

			// Assert
			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			body:           strings.NewReader(`{`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body",
			body: strings.NewReader(`
			{
			"id": "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author": "code learn",
			"title": "first news",
			"summary": "first news post",
			"created_at": "2024-04-07T05:13:27+00:00",
			"source": "https://example.com"
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "db error",
			body: strings.NewReader(`
			{
			"id": "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author": "code learn",
			"title": "first news",
			"summary": "first news post",
			"created_at": "2024-04-07T05:13:27+00:00",
			"source": "https://example.com",
			}`),
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "success",
			body: strings.NewReader(`{
			"id": "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author": "code learn",
			"title": "first news",
			"summary": "first news post",
			"created_at": "2024-04-07T05:13:27+00:00",
			"source": "https://example.com",
			"tags": ["politics"]
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusCreated,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", testCase.body)

			// Act
			handler.PostNews(testCase.store)(w, r)

			// Assert
			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		newsID         string
		expectedStatus int
	}{
		{
			name:           "invalid news id",
			store:          mockNewsStore{},
			newsID:         "invalid-uuid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			store:          mockNewsStore{errState: true},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r.SetPathValue("news_id", testCase.newsID)

			// Act
			handler.GetNewsByID(testCase.store)(w, r)

			// Assert
			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		body           io.Reader
		store          handler.NewsStorer
		expectedStatus int
	}{
		{
			name:           "invalid request body json",
			body:           strings.NewReader(`{`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request body",
			body: strings.NewReader(`
			{
			"id": "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author": "code learn",
			"title": "first news",
			"summary": "first news post",
			"created_at": "2024-04-07T05:13:27+00:00",
			"source": "https://example.com"
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "db error",
			body: strings.NewReader(`
			{
			"id": "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author": "code learn",
			"title": "first news",
			"summary": "first news post",
			"created_at": "2024-04-07T05:13:27+00:00",
			"source": "https://example.com",
			}`),
			store:          mockNewsStore{errState: true},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "success",
			body: strings.NewReader(`{
			"id": "3b082d9d-1dc7-4d1f-907e-50d449a03d45",
			"author": "code learn",
			"title": "first news",
			"summary": "first news post",
			"created_at": "2024-04-07T05:13:27+00:00",
			"source": "https://example.com",
			"tags": ["politics"]
			}`),
			store:          mockNewsStore{},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPut, "/", nil)

			// Act
			handler.UpdateNewsByID(testCase.store)(w, r)

			// Assert
			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		store          handler.NewsStorer
		newsID         string
		expectedStatus int
	}{
		{
			name:           "invalid news id",
			store:          mockNewsStore{},
			newsID:         "invalid-uuid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "db error",
			store:          mockNewsStore{errState: true},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "success",
			store:          mockNewsStore{},
			newsID:         uuid.NewString(),
			expectedStatus: http.StatusNoContent,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, "/", nil)
			r.SetPathValue("news_id", testCase.newsID)

			// Act
			handler.DeleteNewsByID(testCase.store)(w, r)

			// Assert
			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

type mockNewsStore struct {
	errState bool
}

func (m mockNewsStore) Create(handler.NewsPostRequestBody) (news handler.NewsPostRequestBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindByID(_ uuid.UUID) (news handler.NewsPostRequestBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) FindAll() (news []handler.NewsPostRequestBody, err error) {
	if m.errState {
		return news, errors.New("some error")
	}
	return news, nil
}

func (m mockNewsStore) DeleteByID(_ uuid.UUID) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}

func (m mockNewsStore) UpdateByID(_ handler.NewsPostRequestBody) error {
	if m.errState {
		return errors.New("some error")
	}
	return nil
}
