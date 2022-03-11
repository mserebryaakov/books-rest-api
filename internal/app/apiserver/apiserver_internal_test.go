package apiserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)

	if strings.Compare(rec.Body.String(), "Hello rw") != 0 {
		t.Fatal("request no equal")
	}
}
