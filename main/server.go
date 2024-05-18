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
		Title:       "Ascii Art Web",
		Description: "Convert text to ASCII art.",
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
