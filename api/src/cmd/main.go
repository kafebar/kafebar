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

	tlsKeyFile  = os.Getenv("TLS_KEY_FILE")
	tlsCertFile = os.Getenv("TLS_CERT_FILE")
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

	fmt.Printf("starting server on %s\n", addr)

	var err error
	if tlsCertFile != "" {
		err = http.ListenAndServeTLS(addr, tlsCertFile, tlsKeyFile, mux)
	} else {
		err = http.ListenAndServe(addr, mux)
	}

	if err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
