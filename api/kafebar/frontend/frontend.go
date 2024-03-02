package frontend

import (
	"fmt"
	"net/http"
	"net/url"
)

func NewHandler(path string) http.Handler {
	mux := http.NewServeMux()

	indexPath, err := url.JoinPath(path, "index.html")
	if err != nil {
		panic(fmt.Errorf("cannot get index path: %w", err))
	}

	mux.Handle("/", NoCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, indexPath)
	})))

	assetsPath, err := url.JoinPath(path, "assets")
	if err != nil {
		panic(fmt.Errorf("cannot get assets path: %w", err))
	}
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir(assetsPath))))

	return mux
}
