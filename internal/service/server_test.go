package service

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndexHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		statusCode int
		resp       any
	}{
		{
			name:       "found exact value",
			url:        "/endpoint/100",
			statusCode: http.StatusOK,
			resp: SuccessResponse{
				Target:  100,
				Index:   1,
				Value:   100,
				Message: "Value found",
			},
		},
		{
			name:       "found only value with 10%% tolerance",
			url:        "/endpoint/105",
			statusCode: http.StatusOK,
			resp: SuccessResponse{
				Target:  105,
				Index:   1,
				Value:   100,
				Message: "Value found",
			},
		},
		{
			name:       "no value found",
			url:        "/endpoint/1500",
			statusCode: http.StatusNotFound,
			resp: ErrorResponse{
				Message: "Value not found: 1500",
			},
		},
		{
			name:       "bad value type",
			url:        "/endpoint/x",
			statusCode: http.StatusBadRequest,
			resp: ErrorResponse{
				Message: "Bad input value: 'x'",
			},
		},
	}
	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.url, nil)
		w := httptest.NewRecorder()
		srv, err := NewServer(slog.Default(), "./test_data/input.txt")
		if err != nil {
			t.Fatal("cannot load test data")
		}
		srv.router.ServeHTTP(w, req)
		assert.Equal(t, tt.statusCode, w.Code)
		expJson, err := json.Marshal(tt.resp)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		assert.JSONEq(t, string(expJson), w.Body.String())
	}
}

func TestRunServerError(t *testing.T) {
	srv, err := NewServer(slog.Default(), "/bad/path/file.txt")
	assert.ErrorContains(t, err, "cannot load data")
	assert.Len(t, srv.data, 0)

}
