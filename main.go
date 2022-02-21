package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:leo123456@/snippetbox?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS testdb.hello(world varchar(50))")
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec(
		`
	INSERT INTO testdb.hello(world) VALUES ('hello world')
	`)
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT * FROM testdb.hello")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found row containing %q", s)
	}
}
