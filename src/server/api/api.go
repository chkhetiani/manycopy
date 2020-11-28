package api

import (
	"fmt"
	"manycopy/src/server"
	"manycopy/src/server/api/paste-api"
	"manycopy/src/settings"
	"net/http"
	"path/filepath"
	"strings"
)

var s = server.Server {Name: "api" }

func init() {
	for key, val := range pasteapi.Handlers {
		s.AddHandlerFunc(key, val)
	}
	http.HandleFunc("/", fileServerHandler)

	s.Init()
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	// default to index.html
	if r.URL.Path == "/" {
		r.URL.Path = "/index"
	}

	// default extension
	ext := ".html"
	// if has extension don't add default
	if strings.Contains(r.URL.Path, ".") { ext = "" }

	requestedPath := strings.TrimLeft(filepath.Clean(r.URL.Path), "/")
	filename := fmt.Sprintf("%s/%s" + ext, http.Dir(settings.Get("static_path").(string)), requestedPath)
	http.ServeFile(w, r, filename)
}