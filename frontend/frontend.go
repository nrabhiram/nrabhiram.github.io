package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist/static
var embeddedFiles embed.FS

// Get the subtree of the embedded files with `dist` directory as a root.
func BuildHTTPFS() http.FileSystem {
	build, err := fs.Sub(embeddedFiles, "dist/static")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(build)
}
