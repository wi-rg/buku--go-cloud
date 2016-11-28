package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/web1")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping database unsuccessful, we may not use db connection")
	} else {
		fmt.Println("Ping database successful, we may use db connection")
	}

	return db
}

func showPerson(w http.ResponseWriter, r *http.Request) {

	var (
		id int
		name string
	)

	var db = connect()
	defer db.Close()

	rows, err := db.Query("select * from person;")
	if err != nil {
		log.Fatal("Err: ", err)
	}
	defer rows.Close()

	fmt.Fprintln(w, "<table border=1>")
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, "<tr><td>", id, "</td><td>", name, "</td></tr>" )
	}
	fmt.Fprintln(w, "</table>")
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	http.HandleFunc("/", showPerson)       // set router
	errServer := http.ListenAndServe(":9090", nil) // set listen port
	if errServer != nil {
		log.Fatal("ListenAndServe: ", errServer)
	}

}
