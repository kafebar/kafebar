package main

import (
	"fmt"
	"log"
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

	mux.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	addr := fmt.Sprintf(":%s", port)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
