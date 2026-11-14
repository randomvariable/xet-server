package proxyhub

import (
	"net/http/httptest"
	"testing"
)

// A TLS-terminating reverse proxy in front of xet-proxyd forwards plain
// HTTP with X-Forwarded-Proto: https. The refresh route must keep https,
// or hf_xet dials port 80 on the proxy's hostname and hangs.
func TestRefreshRouteURLHonorsForwardedProto(t *testing.T) {
	for _, tc := range []struct {
		fwd, want string
	}{
		{"", "http://hub.example/api/models/o/r/xet-read-token/main"},
		{"https", "https://hub.example/api/models/o/r/xet-read-token/main"},
		{"gopher", "http://hub.example/api/models/o/r/xet-read-token/main"},
	} {
		r := httptest.NewRequest("HEAD", "http://hub.example/o/r/resolve/main/f", nil)
		if tc.fwd != "" {
			r.Header.Set("X-Forwarded-Proto", tc.fwd)
		}
		if got := refreshRouteURL(r, "model", "o/r", "main"); got != tc.want {
			t.Errorf("X-Forwarded-Proto %q: got %q, want %q", tc.fwd, got, tc.want)
		}
	}
}
