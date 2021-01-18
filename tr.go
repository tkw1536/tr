// Package tr provides facilities for a tiny redirect server.
package tr

import "net/http"

// Redirect implements a redirect server that redirects all requests
// See also built-in "http".Redirect
type Redirect struct {
	// Target is the target URL to redirect to
	Target string

	// Permanent determines if the redirect is permanent or temporary
	// The default is temporary
	Permanent bool
}

// ServeHTTP implements the http.Handler interface and redirects a single request to redirect.Target
func (redirect Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dest := redirect.Target + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		dest += "?" + r.URL.RawQuery
	}

	status := http.StatusTemporaryRedirect
	if redirect.Permanent {
		status = http.StatusPermanentRedirect
	}

	http.Redirect(w, r, dest, status)
}

func init() {
	var _ http.Handler = (*Redirect)(nil)
}
