// Package tr contains a tiny redirect server implementation
package tr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect_ServeHTTP(t *testing.T) {
	type fields struct {
		Target    string
		Permanent bool
		Overrides map[string]string
	}
	tests := []struct {
		name         string
		fields       fields
		path         string
		wantCode     int
		wantLocation string
	}{
		{"non-permanent redirect /", fields{"https://example.com/subdirectory", false, nil}, "/", http.StatusTemporaryRedirect, "https://example.com/subdirectory/"},
		{"permanent redirect /", fields{"https://example.com/subdirectory", true, nil}, "/", http.StatusPermanentRedirect, "https://example.com/subdirectory/"},

		{"non-permanent redirect /path/", fields{"https://example.com/subdirectory", false, nil}, "/path/", http.StatusTemporaryRedirect, "https://example.com/subdirectory/path/"},
		{"permanent redirect /path/", fields{"https://example.com/subdirectory", true, nil}, "/path/", http.StatusPermanentRedirect, "https://example.com/subdirectory/path/"},

		{"non-permanent overwrite redirect", fields{"https://example.com/subdirectory", false, map[string]string{"/overwrite": "https://overwrite/"}}, "/overwrite", http.StatusTemporaryRedirect, "https://overwrite/"},
		{"non-permanent overwrite redirect", fields{"https://example.com/subdirectory", true, map[string]string{"/overwrite": "https://overwrite/"}}, "/overwrite", http.StatusPermanentRedirect, "https://overwrite/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redirect := Redirect{
				Target:    tt.fields.Target,
				Permanent: tt.fields.Permanent,
				Overrides: tt.fields.Overrides,
			}

			// make a test request
			req := httptest.NewRequest("", "http://localhost:8080"+tt.path, nil)
			w := httptest.NewRecorder()
			redirect.ServeHTTP(w, req)

			// check that the code is right
			if w.Code != tt.wantCode {
				t.Errorf("redirect.ServeHTTP(): got code = %d, want %d", w.Code, tt.wantCode)
			}

			// check that the location header is set right
			gotLocations, _ := w.HeaderMap["Location"]
			var gotLocation string
			if len(gotLocations) == 1 {
				gotLocation = gotLocations[0]
			}
			if gotLocation != tt.wantLocation {
				t.Errorf("redirect.ServeHTTP(): got location = %s, want %s", gotLocation, tt.wantLocation)
			}
		})
	}
}
