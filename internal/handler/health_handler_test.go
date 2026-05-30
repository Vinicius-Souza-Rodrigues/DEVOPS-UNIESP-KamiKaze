package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(
	t *testing.T,
) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rec := httptest.NewRecorder()

	Health(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected %d got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}
