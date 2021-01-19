// Package tr implements a small
package tr

import (
	"net/http"
	"strings"
)

// Redirect implements a redirect server that redirects all requests.
// It implements http.Handler.
type Redirect struct {
	// Target is the target URL to redirect to.
	Target string

	// Absolute determines if the request path should be appended to the target URL when redirecting.
	// By default this path is always appended, set Absolute to true to prevent this.
	Absolute bool

	// Overrides is a map from paths to URLs that should override the default target.
	Overrides map[string]string

	// Permanent determines if the redirect responses issued should return
	// Permanent Redirect (Status Code 308) or Temporary Redirect (Status Code 307).
	Permanent bool
}

// Redirect determines the redirect URL for a specific incoming request
func (redirect Redirect) Redirect(r *http.Request) string {
	// if we have an override for this URL, use it immediatly
	url := strings.TrimSuffix(r.URL.Path, "/")
	if override, ok := redirect.Overrides[url]; ok {
		return override
	}

	// if we are in absolute redirect mode, always return the absolute URL
	if redirect.Absolute {
		return redirect.Target
	}

	// return the target + the redirected URL
	dest := strings.TrimSuffix(redirect.Target, "/") + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		dest += "?" + r.URL.RawQuery
	}
	return dest
}

// ServeHTTP implements the http.Handler interface and redirects a single request to redirect.Target.
func (redirect Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dest := redirect.Redirect(r)

	// determine if we are temporary or permanent redirect
	status := http.StatusTemporaryRedirect
	if redirect.Permanent {
		status = http.StatusPermanentRedirect
	}

	// and do the redirect
	http.Redirect(w, r, dest, status)
}

func init() {
	var _ http.Handler = (*Redirect)(nil)
}
