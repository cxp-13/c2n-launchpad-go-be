package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseInfoEndpoint(t *testing.T) {
	r := SetupGinRouter()
	testCases := []struct {
		name       string
		url        string
		statusCode int
	}{
		{
			name:       "Valid product ID",
			url:        "/product/base_info?productId=1",
			statusCode: http.StatusOK,
		},
		{
			name:       "Invalid product ID",
			url:        "/product/base_info?productId=0",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Missing product ID",
			url:        "/product/base_info",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.url, nil)
			assert.NoError(t, err)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestListEndpoint(t *testing.T) {
	r := SetupGinRouter()
	req, err := http.NewRequest(http.MethodGet, "/product/list", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Further assertions on response body or other details can be added here
}
