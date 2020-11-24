package api

import (
	"fmt"
	"manycopy/src/server"
	"manycopy/src/server/api/paste-api"
	"net/http"
)

var s = server.Server {Name: "api" }

func init() {

	for key, val := range pasteapi.Handlers {
		s.AddHandlerFunc(key, val)
	}
	
	s.AddHandlerFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})
	s.Init()
}