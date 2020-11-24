package pasteapi

import (
	"fmt"
	"net/http"
)

var Handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/paste/new": paste,
}

func paste (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "h")
}