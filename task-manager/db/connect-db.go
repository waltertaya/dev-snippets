package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite", "data/tasks.sqlite")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected to the SQLite database successfully.")

	// get the version of the SQLite
	var sqliteVersion string
	err = DB.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sqliteVersion)
}
