package repo

import (
	"database/sql"
	"fmt"
	"log"
)

func database() *sql.DB {
	databaseString := "postgres://username:password@your-database-url:5432/your-database-name"
	db, error := sql.Open("postgres", databaseString)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connected to Database")
	return db
}
