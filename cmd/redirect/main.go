package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tkw1536/tr"
)

func main() {
	fmt.Printf("Redirecting %s to %s\n", bindAddress, redirect.Target)
	http.ListenAndServe(bindAddress, redirect)
}

var bindAddress string
var redirect tr.Redirect

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: TARGET=<TARGET> [PERMANENT=1] redirect <bindAdress>")
		os.Exit(1)
	}
	bindAddress = os.Args[1]
	redirect.Target = os.Getenv("TARGET")
	redirect.Permanent = os.Getenv("PERMANENT") == "1"
}
