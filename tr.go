package main

import (
	"fmt"
	"net/http"
	"os"
)

var redirectTarget, bindAddress string

func serve(w http.ResponseWriter, req *http.Request) {
	dest := redirectTarget + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		dest += "?" + req.URL.RawQuery
	}
	http.Redirect(w, req, dest, http.StatusTemporaryRedirect)
}

func main() {
	// read command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: TARGET=<TARGET> redirect <bindAdress>")
		os.Exit(1)
	}
	bindAddress = os.Args[1]
	redirectTarget = os.Getenv("TARGET")

	fmt.Printf("Redirecting %s to %s\n", bindAddress, redirectTarget)

	// and start an http server
	http.ListenAndServe(bindAddress, http.HandlerFunc(serve))
}
