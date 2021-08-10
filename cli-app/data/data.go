package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func OpenDatabase() error{
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
if err!=nil{
	return err
}
    return  db.Ping()
}

func CreateTable()  {
	CreateTableSql :=  `CREATE TABLE IF NOT EXISTS CLI(
"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
"word" TEXT,
"Definition" TEXT,
"Category" TEXT
);`
 stmt, err := db.Prepare(CreateTableSql)
 if err!=nil{
 	log.Fatalln(err.Error())
 }
 stmt.Exec()
 log.Println("table created")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO studybuddy(word, definition, category) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted study note successfully")
}