package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDefaultRoute(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	setupRouter(r)

	req, err := http.NewRequestWithContext(ctx, "GET", "/", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}

	r.ServeHTTP(w, req)
	if http.StatusOK != w.Code {
		t.Fatalf("expected response code %d, got %d", http.StatusOK, w.Code)
	}

	body := w.Body.String()

	expected := "success"

	if !strings.Contains(body, expected) {
		t.Fatalf("expected body to contain '%s', got %s", expected, body)
	}
} // End
