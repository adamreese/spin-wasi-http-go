package main

import (
	"fmt"
	"net/http"

	"github.com/spinframework/spin-go-sdk/v2/variables"
	"github.com/spinframework/spin-go-sdk/v2/wit"
	"github.com/ydnar/wasi-http-go/wasihttp"
)

var _ = wit.Wit

func init() {
	wasihttp.Serve(&WasiHTTP{})
}

type WasiHTTP struct{}

func (WasiHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Get variable value `message` defined in spin.toml.
	val, err := variables.Get("message")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "message: ", val)

	fmt.Fprintln(w, "Hello World!")
}

func main() {}
