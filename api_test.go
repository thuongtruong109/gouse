package gouse

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpError(t *testing.T) {
	_, err := http.NewRequest("GET", "/img", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	_httpError(rr, "Not Found", http.StatusNotFound)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, status)
	}

	expected := "Not Found"
	actual := strings.TrimSpace(rr.Body.String())

	if actual != expected {
		t.Errorf("Expected body %q, got %q", expected, actual)
	}
}
