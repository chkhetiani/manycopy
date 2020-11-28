package pasteapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chkhetiani/manycopy/src/db"
)

var Handlers = map[string]func(http.ResponseWriter, *http.Request){
	"/paste/new": paste,
	"/paste/get": get,
}

func paste(w http.ResponseWriter, r *http.Request) {
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

	var dbData []db.Input
	for _, inp := range data {
		if (inp.Type == "INPUT" || inp.Type == "TEXTAREA") && inp.Data != "" {
			dbData = append(dbData, inp)
		}
	}

	if len(dbData) == 0 {
		w.WriteHeader(400)
		_, err := fmt.Fprintf(w, "bad inputs")
		if err != nil {
			panic(err)
		}
		return
	}

	id := db.InsertPaste(dbData)

	x, _ := json.Marshal(id)

	_, err = fmt.Fprintf(w, string(x))
	if err != nil {
		log.Fatalln(err.Error())
	}
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

	_, err := fmt.Fprintf(w, string(x))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
