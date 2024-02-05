package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	port   = os.Getenv("PORT")
	uiPath = os.Getenv("UI_PATH")
)

func main() {
	mux := http.NewServeMux()

	if uiPath != "" {
		mux.Handle("/", http.FileServer(http.Dir(uiPath)))
	}

	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, mux)
}
