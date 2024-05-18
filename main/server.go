package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"ascii"
)

type Info struct {
	Title       string
	Description string
	Font1       string
	Font2       string
	Font3       string
	Font        string
	Content     string
	Text        string
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
	text := r.FormValue("text")
	fonts := map[string]string{
		"Font1": "standard",
		"Font2": "shadow",
		"Font3": "thinkertoy",
	}
	Font := r.FormValue("select")

	Data := Info{
		Title:       "Ascii Art Web",
		Description: "Convert text to ASCII art.",
		Font1:       fonts["Font1"],
		Font2:       fonts["Font2"],
		Font3:       fonts["Font3"],
		Text:        text,
		Content:     ascii.Asci(text, ascii.DefaultFont(Font, "standard")),
	}

	err = templ.Execute(w, Data)
	// make the text empty
	Data.Text = ""
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("../output/ascii-art-web.txt", []byte(Data.Content), 0o644)
}

func main() {
	http.HandleFunc("/", pageHandler)
	log.Println("\033[32mServer is running on port 8080...🚀\033[0m")
	log.Println("\033[32mhttp://localhost:8080\033[0m")
	http.ListenAndServe(":8080", nil)
}
