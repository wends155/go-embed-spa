package frontend

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// Embed the build directory from the frontend.
//
//go:embed build/*
//go:embed build/_app/immutable/*
//go:embed build/_app/immutable/assets/*
var BuildFs embed.FS

// Get the subtree of the embedded files with `build` directory as a root.
func BuildHTTPFS() http.FileSystem {
	build, err := fs.Sub(BuildFs, "build")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(build)
}
