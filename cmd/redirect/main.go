// Command redirect implements an HTTP server that answers every incoming request with a redirect.
//
//   TARGET=<TARGET> [ABSOLUTE=1] [PERMANENT=1] [OVERRIDES=/path/to/.json] redirect <bindAddress>
//
// The address to bind to is specified by the bindAdress parameter.
// By default, this server redirects all requests to the URL specified by the TARGET environment variable.
//
// For each incoming request, the request path is append to it.
// For instance, if target is "http://example.com" and the request path is "/index.html", the server will redirect to "http://example.com/index.html".
// To disable this behavior and always redirect to the exact TARGET set the ABSOLUTE environment variable to "1".
//
// Without additional configuration, all redirect responses return HTTP Status Code 307 (Temporary Redirect).
// To instead use Status Code 308 (Permanent Redirect), set the PERMANENT environment variable to "1".
//
// To further override default behavior, the server can also change the behavior for individual request paths.
// For this purpose, set the OVERRIDES environment variable to a filepath containing a .json file.
// This .json file is assumed to contain an object mapping request URLS to target URLS.
// The request URLS are assumed to have trailing '/'s removed.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tkw1536/tr"
)

func main() {
	fmt.Printf("Redirecting %s to %s\n", bindAddress, redirect.Target)
	fmt.Printf("Loaded %d overwrite(s)\n", len(redirect.Overrides))
	http.ListenAndServe(bindAddress, redirect)
}

var bindAddress string
var redirect tr.Redirect

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: TARGET=<TARGET> [ABSOLUTE=1] [PERMANENT=1] [OVERRIDES=/path/to/.json] redirect <bindAddress>")
		os.Exit(1)
	}

	bindAddress = os.Args[1]

	redirect.Target = os.Getenv("TARGET")
	redirect.Absolute = os.Getenv("ABSOLUTE") == "1"
	redirect.Permanent = os.Getenv("PERMANENT") == "1"

	overridesPath := os.Getenv("OVERRIDES")
	if overridesPath != "" {
		bytes, err := ioutil.ReadFile(overridesPath)
		if err != nil {
			panic("Unable to read " + overridesPath + ": " + err.Error())
		}
		if err := json.Unmarshal(bytes, &redirect.Overrides); err != nil {
			panic("Unable to parse " + overridesPath + ": " + err.Error())
		}

	}
}
