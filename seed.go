package main

import (
	// "html/template"
	"net/http"
	"io"
)

func HandleSeed(w http.ResponseWriter, r *http.Request) {
	CreateStudent(student{"Anaëlle", "Unity, RaspPi, Docker, Golang", "4 saisons",
		"Cet été la - Joe Hisaishi", "Anzouz",
		"Remember to always be yourself unless you suck. Then pretend to be someone else"})
	io.WriteString(w, "seed 👌\n")
}