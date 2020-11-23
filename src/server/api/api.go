package api

import (
	"fmt"
	"manycopy/src/server"
	"net/http"
)

func init() {
	ser := server.Server {Name: "api" }
	ser.AddHandlerFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})
	ser.Init()
}