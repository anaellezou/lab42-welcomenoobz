package main

import (
	//"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"html/template"
	"io"
	// "path/filepath"
	//"log"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("home.html")
	fmt.Println(err)
	t.Execute(w, nil)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello")
}

func main() {
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/static", ServeStatic)
	http.ListenAndServe(":4242", nil)
}