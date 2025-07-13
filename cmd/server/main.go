package main

import (
	"fmt"
	"net/http"
	"os"
	"sapphire/frontend"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Get the filesystem
	staticFS := frontend.BuildHTTPFS()
	fileServer := http.FileServer(staticFS)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path

		htmlPath := urlPath

		if htmlPath == "/" {
			htmlPath = "index.html"
		} else if !strings.HasSuffix(htmlPath, ".html") {
			htmlPath += ".html"
		}

		fmt.Println(htmlPath)

		f, err := staticFS.Open(htmlPath)
		if err == nil {
			// HTML file exists
			f.Close()
			fmt.Println("FOUND")
			// Rewrite to serve the HTML file
			if htmlPath != "index.html" {
				r.URL.Path = "/" + htmlPath
			}
		} else if !strings.Contains(urlPath, ".") {
			// If nothing else matches, serve index.html for client-side routing
			r.URL.Path = "/"
		}

		fileServer.ServeHTTP(w, r)
	})

	println("Server running on :" + port)
	http.ListenAndServe(":"+port, nil)
}
