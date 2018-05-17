package main

import (
	//"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"html/template"
	//"log"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	fmt.Println(err)
	t.Execute(w, nil)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	w.writeHeader(http.StatusOk)
	io.WriteString(w, "Hello")
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/static", ServeStatic)
	http.ListenAndServe(":4242", nil)
}