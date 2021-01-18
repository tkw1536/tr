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
		fmt.Println("Usage: TARGET=<TARGET> [PERMANENT=1] [OVERRIDES=/path/to/.json] redirect <bindAdress>")
		os.Exit(1)
	}

	bindAddress = os.Args[1]
	redirect.Target = os.Getenv("TARGET")
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
