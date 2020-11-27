package api

import (
	"fmt"
	"log"
	"manycopy/src/server"
	"manycopy/src/server/api/paste-api"
	"manycopy/src/settings"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var s = server.Server {Name: "api" }

func init() {

	for key, val := range pasteapi.Handlers {
		s.AddHandlerFunc(key, val)
	}

	http.HandleFunc("/", x)

	//fserver  := http.FileServer(http.Dir(settings.Get("static_path").(string)))
	//s.AddHandler("/", fserver)
	s.Init()

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
}

func x(w http.ResponseWriter, r *http.Request) {
	ext := ".html"
	if r.URL.Path == "/" {
		r.URL.Path = "/index"
	}

	if strings.Contains(r.URL.Path, ".") { ext = "" }

	requestedPath := strings.TrimLeft(filepath.Clean(r.URL.Path), "/")
	filename := fmt.Sprintf("%s/%s" + ext, http.Dir(settings.Get("static_path").(string)), requestedPath)
	http.ServeFile(w, r, filename)
}