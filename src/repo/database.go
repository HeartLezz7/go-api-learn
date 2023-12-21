package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	databaseString := "root:123456@tcp(localhost:3306)/cc15_todo_list_v2"
	db, error := sql.Open("mysql", databaseString)
	if error != nil {
		fmt.Println("SERVER ERROR", error.Error())
	}
	error = db.Ping()
	if error != nil {
		fmt.Println("SERVER ERROR", error.Error())
	}
	fmt.Println("Connected to Database")
	return db
}
