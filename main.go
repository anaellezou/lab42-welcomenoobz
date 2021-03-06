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

type ViewData struct {
	Studs []student
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	stud, err := GetAllStudents()
	t, err := template.ParseFiles("home.html")
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
	}
	viewdata := ViewData{stud}
    // fmt.Println(view)
	err = t.Execute(w, viewdata)
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
	stud, _ := FindStudentBy("name", name)
	response.Reply("> *" + stud.Name + "* mange des pizzas *" + stud.Pizza + "* en écoutant *" + stud.Music + "* en codant du *" + stud.Techs + "*")
}

func handleNoob(request slacker.Request, response slacker.ResponseWriter) {
	response.Typing()
	name := request.Param("name")
	techs := request.Param("techs")
	pizza := request.Param("pizza")
	music := request.Param("music")
	role := request.Param("role")
	inspirationnal_quote := request.Param("inspirationnal_quote")
	image := request.Param("image")
	if name == "" || techs == "" || pizza == "" || music == "" || role == "" || inspirationnal_quote == "" {
		response.Reply("> Usage: @gossiplab index <name> <pizza> <fav_song> <role> <inspirationnal_quote> <image>")
		return
	} 
	CreateStudent(student{name, techs, pizza, music, role, inspirationnal_quote, image})
	response.Reply("> Noob has been created in db")
}

func main() {

	TestDB()
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
	bot.Command("index <name> <techs> <pizza> <music> <role> <inspirationnal_quote>", "Index a noob of the lab",
		handleNoob)
	go bot.Listen()

	// }()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/seed", HandleSeed)

	// Server start
	fmt.Printf("Server started.\n")
	fmt.Printf("Error: %v\n", http.ListenAndServe(":4242", nil))
}
