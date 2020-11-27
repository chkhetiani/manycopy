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
		port := os.Getenv("PORT")

		if port == "" {
			port = settings.Get(s.Name + "_port").(string)
		} else {
			port = "0.0.0.0:" + port
		}

		fmt.Printf("starting %v server at %v ...\n", s.Name, port)

		if err := http.ListenAndServe(port, nil); err != nil {
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