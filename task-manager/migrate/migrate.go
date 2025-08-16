package main

import (
	"fmt"
	"log"

	"github.com/waltertaya/dev-snippets/task-manager/db"
	"github.com/waltertaya/dev-snippets/task-manager/initialisers"
)

func init() {
	initialisers.LoadEnv()
	db.ConnectDB()
}

func main() {
	_, err := db.DB.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, done BOOLEAN NOT NULL)")
	if err != nil {
		log.Fatalf("Error creating table: %q\n", err)
	}
	fmt.Println("Auto-migrated successfully")
}
