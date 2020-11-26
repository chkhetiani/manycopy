package pasteapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"manycopy/src/db"
	"net/http"
)

var Handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/paste/new": paste,
	"/paste/get": get,
}

func paste(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "POST" {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data []db.Input
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	id := db.InsertPaste(data)

	x, _ := json.Marshal(id)

	fmt.Fprintf(w, string(x))
}



func get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	id := ids[0]

	z := db.GetPaste(id)

	x, _ := json.Marshal(z)

	fmt.Fprintf(w, string(x))
}