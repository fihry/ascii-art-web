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
	Text        string
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	templ, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	// fonts := map[string]string {
	// 	"Font1": "standard",
	// 	"Font2": "shadow",
	// 	"Font3": "thinkertoy",
	// }
	text := r.FormValue("text")
	Data := Info{
		Title:       "AsciiArtWeb",
		Description: "Make your Art",
		Font1:       "standard",
		Font2:       "shadow",
		Font3:       "thinkertoy",
		Text:        text,
		Content:     ascii.Asci(text, "standard"),
	}
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	err = templ.Execute(w, Data)
	Data.Text = ""
	if err != nil {
		log.Fatal(err)
	}
}

// func AsciiArtWebHundler (w http.ResponseWriter,re *http.Request){
// 	if
// }

func main() {
	http.HandleFunc("/", pageHandler)
	// http.Handle("/ascii-art", nil)
	http.ListenAndServe(":8080", nil)
}
