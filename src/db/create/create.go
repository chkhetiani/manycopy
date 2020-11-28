package create

import (
	"database/sql"
	"log"
	"os"
)

var FileName = "./db.db"

func init() {
	_, err := os.Stat(FileName)
	exists := !os.IsNotExist(err)
	if !exists {
		file, err := os.Create(FileName)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = file.Close()
		if err != nil {
			panic(err)
		}

		db, err := sql.Open("sqlite3", FileName)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := db.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		createTables(db)
	}
}

func createTables(db *sql.DB) {
	createInputsTable(db)
	createPastesTable(db)
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
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}
