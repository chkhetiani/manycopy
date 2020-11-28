package db

import (
	"database/sql"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/chkhetiani/manycopy/src/db/create"
	_ "github.com/mattn/go-sqlite3"
)

var pasteIdLength = 6

type Input struct {
	Data string
	Type string
}

func getId() string {
	rand.Seed(time.Now().Unix())

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var id strings.Builder
	for i := 0; i < pasteIdLength; i++ {
		id.WriteString(string(charset[rand.Intn(52)]))
	}

	return id.String()
}

func InsertPaste(data []Input) string {
	id := getId()
	query := `INSERT INTO pastes(id) VALUES(?)`

	db, err := sql.Open("sqlite3", create.FileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatalln(err.Error())
	}

	query = `INSERT INTO inputs(data, type, pasteId) VALUES(?, ?, ?)`
	statement, err = db.Prepare(query)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, inp := range data {
		_, err = statement.Exec(inp.Data, inp.Type, id)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	return id
}

func GetPaste(id string) []Input {
	db, err := sql.Open("sqlite3", create.FileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	row, err := db.Query("SELECT data, type FROM inputs WHERE pasteId = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = row.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()
	var res []Input
	for row.Next() {
		var input Input
		if err := row.Scan(&input.Data, &input.Type); err != nil {
			log.Fatal(err)
		}
		res = append(res, input)
	}

	return res
}
