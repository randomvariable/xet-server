package casserver

import (
	"net/http/httptest"
	"testing"
)

// Xorb fetch URLs in reconstruction responses are built from the request's
// scheme; behind a TLS-terminating reverse proxy that must come from
// X-Forwarded-Proto, or clients fetch xorbs over plain http.
func TestBaseURLFromRequestHonorsForwardedProto(t *testing.T) {
	r := httptest.NewRequest("GET", "http://cas.example/v1/reconstructions/x", nil)
	if got := baseURLFromRequest(r); got != "http://cas.example" {
		t.Errorf("no header: got %q", got)
	}
	r.Header.Set("X-Forwarded-Proto", "https")
	if got := baseURLFromRequest(r); got != "https://cas.example" {
		t.Errorf("X-Forwarded-Proto https: got %q", got)
	}
}
