package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	// "log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shomali11/slacker"
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

func handleBot(request slacker.Request, response slacker.ResponseWriter) {
	name := request.Param("name")
	if name == "" {
		response.Reply("> Usage: @hellobot hello Name")
		return
	}
	response.Reply("Hey " + name + "!")
}

func handleWhoIs(request slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	name := request.Param("name")
	if name == "" {
		response.Reply("> Usage: @gossiplab whois <name>")
		return
	}
	stud, _ := FindStudentBy("name", "Thomas")
	response.Reply("> *" + stud.name + "* mange des pizzas *" + stud.pizza + "* en écoutant *" + stud.music + "* en codant du *" + stud.techs + "*")
}

func handleNoob(request slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	name := request.Param("name")
	techs := request.Param("techs")
	pizza := request.Param("pizza")
	music := request.Param("music")
	role := request.Param("role")
	inspirationnal_quote := request.Param("inspirationnal_quote")
	if name == "" || techs == "" || pizza == "" || music == "" || role == "" || inspirationnal_quote == "" {
		response.Reply("> Usage: @gossiplab index <name> <pizza> <fav_song> <role> <inspirationnal_quote>")
		return
	} 
	CreateStudent(student{name, techs, pizza, music, role, inspirationnal_quote})
	response.Reply("> Noob has been created in db")
}

func main() {
	TestDB()
	thom := student{"Thomas", "GOlang", "hawaïenne", "pf:wishuwerehere", "racine", "wsh"}
	CreateStudent(thom)

	// Launching bot
	bot := slacker.NewClient(os.Getenv("API_TOKEN"))
	bot.Init(func() {
		fmt.Printf("Connected!")
	})
	bot.Err(func(err string) {
		fmt.Printf(err)
	})
	bot.DefaultCommand(func(request slacker.Request, response slacker.ResponseWriter) {
		response.Reply("Say what?")
	})
	bot.Command("hello <name>", "Say hello to someone", handleBot)
	bot.Command("whois <name>", "Ask information about someone", handleWhoIs)
	bot.Command("index <name> <pizza> <fav_song> <role> <inspirationnal_quote>", "Index a noob of the lab",
		handleNoob)
	bot.Listen()

	// }()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", HomePage)

	// Server start
	fmt.Printf("Server started.\n")
	fmt.Printf("Error: %v\n", http.ListenAndServe(":4242", nil))
}
