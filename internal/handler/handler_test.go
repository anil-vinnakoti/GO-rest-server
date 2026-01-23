package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anil-vinnakoti/newsapi/internal/handler"
)

func Test_GetAllNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{name: "not implemented",
			expectedStatus: http.StatusNotImplemented},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			handler.GetAllNews()(w, r)

			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_PostNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{name: "not implemented",
			expectedStatus: http.StatusNotImplemented},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			handler.PostNews()(w, r)

			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{name: "not implemented",
			expectedStatus: http.StatusNotImplemented},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			handler.GetNewsByID()(w, r)

			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{name: "not implemented",
			expectedStatus: http.StatusNotImplemented},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPut, "/", nil)

			handler.UpdateNewsByID()(w, r)

			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_DeleteNewsByID(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{name: "not implemented",
			expectedStatus: http.StatusNotImplemented},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodDelete, "/", nil)

			handler.DeleteNewsByID()(w, r)

			if w.Result().StatusCode != testCase.expectedStatus {
				t.Errorf("expected: %d, recieved: %d", testCase.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
