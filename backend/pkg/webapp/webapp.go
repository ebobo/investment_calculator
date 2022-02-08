package webapp

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
)

//go:embed dist/*

// WebAppFS is a read-only collection of Vue web application files
var WebAppFS embed.FS

// SpaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type SpaHandler struct {
	// The path to prepend the request when trying to access the bundled webapp
	StaticPath string
	// The frontend bundle as an go embed pragma
	Bundle embed.FS
}

// ServeHTTP inspects the URL path to locate a file within the frontend bundle
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		fmt.Println("Failed to fetch path from request.")
		fmt.Println(err)
		return
	}
	fmt.Printf("SpaHandler: Request path '%s'", path)

	sub, err := fs.Sub(h.Bundle, "dist")
	if err != nil {
		fmt.Println("Failed to open bundled frontend.")
		return
	}
	httpFs := http.FS(sub)

	_, err = h.Bundle.Open(h.StaticPath + path)

	if err != nil {
		// The file does not exist in the webapp bundle. We now want to check whether
		// the request was a specific file, in which we want to let it pass through. If
		// the request is a path (ie no ext), we serve index.html by putting a cloned req
		// with the / as a path.
		fileExt := filepath.Ext(path)
		if fileExt == "" {
			fmt.Println("SpaHandler: Tried to access non existing path. Sending index.")
			clonedReq := r.Clone(r.Context())
			if path != "/" {
				clonedReq.URL.Path = "/"
			}

			http.FileServer(httpFs).ServeHTTP(w, clonedReq)
			return
		}

		fmt.Printf("SpaHandler: Tried to access a non existant file in frontend bundle '%s'", path)
	}

	http.FileServer(httpFs).ServeHTTP(w, r)
}
