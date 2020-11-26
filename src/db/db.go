package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var pasteIdLength = 6
var fileName = "./db.db"

type Input struct {
	Data string
	Type string
}


func init() {
	_, err := os.Stat(fileName)
	exists := !os.IsNotExist(err)
	fmt.Println(exists)
	if !exists {
		fmt.Println("not exists")
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()

		db, err := sql.Open("sqlite3", fileName)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		createTables(db)
	}
}

func createInputsTable(db *sql.DB) {
	createSQL := `CREATE TABLE inputs (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"data" TEXT,
		"type" integer,
		"pasteId" TEXT references pastes(id)
	  );`

	statement, err := db.Prepare(createSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createPastesTable(db *sql.DB) {
	createSQL := `CREATE TABLE pastes (
		"id" TEXT NOT NULL PRIMARY KEY		
	  );`

	statement, err := db.Prepare(createSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func createTables(db *sql.DB) {
	createInputsTable(db)
	createPastesTable(db)
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
	query:= `INSERT INTO pastes(id) VALUES(?)`

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()


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
	if err != nil { log.Fatalln(err.Error()) }

	for _, inp := range data {
		_, err = statement.Exec(inp.Data, inp.Type, id)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	return id
}

func GetPaste(id string) []Input {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT data, type FROM inputs WHERE pasteId = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
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
