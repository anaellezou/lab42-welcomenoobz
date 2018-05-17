package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}
}

func testDB() {
	db, err := sql.Open("sqlite3", "./lab42.db")
	if err != nil {
		fmt.Printf("Error DB: %v\n", err)
	}
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Printf("%v\n", db)
	defer db.Close()
}

func main() {
	testDB()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", HomePage)

	// Server start
	fmt.Printf("Server started.\n")
	fmt.Printf("Error: %v\n", http.ListenAndServe(":4242", nil))
}
