package main

import (
	"html/template"
	"log"
	"net/http"

	"ascii"
)

type Info struct {
	Title       string
	Description string
	Font1       string
	Font2       string
	Font3       string
	Content     string
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templ, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fonts := map[string]string {
		"Font1": "standard",
		"Font2": "shadow",
		"Font3": "thinkertoy",
	}
	Data := Info{
		Title:       "AsciiArtWeb",
		Description: "Make your Art",
		Font1:       "standard",
		Font2:       "shadow",
		Font3:       "thinkertoy",
		Content:     "",
	}
	var text, font string
	if r.Method == "POST" {
		text = r.FormValue("text")
		font = r.FormValue("font")
		log.Print(text, "standard")
	}
	Data.Content = ascii.Asci(text, fonts[font])
	err = templ.Execute(w, Data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe(":8080", nil)
}
