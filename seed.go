package main

import (
	// "html/template"
	"net/http"
	"io"
)

func HandleSeed(w http.ResponseWriter, r *http.Request) {
	CreateStudent(student{"Anaëlle", "Unity, RaspPi, Docker, Golang", "4 saisons",
		"Cet été la - Joe Hisaishi", "Anzouz",
		"Remember to always be yourself unless you suck. Then pretend to be someone else", "ananelle.png"})
	CreateStudent(student{"Clément", "Docker, Flask, Python, Rancher, Git", "4 fromages",
		"Une Matina - Ludovico Einaudi", "Malefoy, le rappeur",
		"Always shoot for the moon, even if you miss you'll land among the stars", "cmalfroy.png"})
	CreateStudent(student{"Thomas", "Ruby on Rails, git, docker, bash", "Marguarita avec Origan",
		"Canon in D - Pachelbel", "Racine",
		"Quote de porc", "tpayet.jpg"})
	CreateStudent(student{"Qarc", "Go, Kafka, Nodejs, Meshlab, Slack, Esport (vs-fighting)",
		"Chèvre-miel", "Sans-repères - Sniper", "Qarciflette, le ldap", "L'homme est un roseau pensant", "qarc.jpg"})
	CreateStudent(student{"Marco", "Swift, Ruby on rails, xCode, Raspberry", "La Reine - Si elle est dégueu le resto est pas ouf",
		"A day in the life - The Beatles", "En couple avec le root", "Inutile donc indispensable", "marco.png"})
	CreateStudent(student{"Christophe", "PHP Laravel, Python, JS, jQuery, iOS, git", "Caleçon",
		"Ding ding dong - Gunther", "Le maquereau", "Ne pas oublier les autres 50%", "oseng.jpg"})
	io.WriteString(w, "seed 👌\n")
}