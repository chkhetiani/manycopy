package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/chkhetiani/manycopy/src/server"
	pasteapi "github.com/chkhetiani/manycopy/src/server/api/paste-api"
	"github.com/chkhetiani/manycopy/src/settings"
)

var s = server.Server{Name: "api"}

func init() {
	for key, val := range pasteapi.Handlers {
		s.AddHandlerFunc(key, val)
	}
	http.HandleFunc("/", fileServerHandler)

	s.Init()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	// default to index.html
	if r.URL.Path == "/" {
		r.URL.Path = "/index"
	}

	// default extension
	ext := ".html"
	// if has extension don't add default
	if strings.Contains(r.URL.Path, ".") {
		ext = ""
	}

	requestedPath := strings.TrimLeft(filepath.Clean(r.URL.Path), "/")
	filename := fmt.Sprintf("%s/%s"+ext, http.Dir(settings.Get("static_path").(string)), requestedPath)
	fmt.Println(r.RemoteAddr + ": " + filename)
	http.ServeFile(w, r, filename)
}
