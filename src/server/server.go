package server

import (
	"fmt"
	"log"
	"manycopy/src/settings"
	"net/http"
	"os"
)

type Server struct {
	Name string
}

func (s *Server) Init() {
	go func(s *Server) {
		hostname := getHostName(s)
		fmt.Printf("starting at %v", hostname)
		if err := http.ListenAndServe(hostname, nil); err != nil {
			log.Fatal(err)
		}
	}(s)
}

func (s *Server) AddHandlerFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, f)
}

func (s *Server) AddHandler(path string, h http.Handler) {
	http.Handle(path, h)
}

func getHostName(s *Server) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = settings.Get(s.Name + "_port").(string)
	}
	return "0.0.0.0" + port
}