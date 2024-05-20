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
	Font        string
	Content     string
	Text        string
	Warning     string
}

var Data = Info{
	Title:       "Ascii Art Web",
	Description: "Convert text to ASCII art.",
	Text:        "",
	Content:     "",
	Warning:     "",
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

	err = templ.Execute(w, Data)
	if err != nil {
		log.Fatal(w, "Could not execute template", http.StatusInternalServerError)
	}
	Data.Text, Data.Content, Data.Warning = "", "", ""
}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		log.Println("Could not parse form")
	}

	// Get the value of a form field and assign it to the Data struct
	Data.Text = r.FormValue("text")
	if !ascii.IsPrintable(Data.Text) {
		Data.Warning = "The text contains non-printable characters."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	Data.Font = r.FormValue("select")
	if Data.Font == "" {
		Data.Font = "standard"
	}
	Data.Content = ascii.AsciiArt(Data.Text, Data.Font)
	// Then redirect the user to the root page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/ascii-art", HandleAscii)
	log.Println("\033[32mServer is running on port 8080...🚀\033[0m")
	log.Println("\033[32mhttp://localhost:8080\033[0m")
	http.ListenAndServe(":8080", nil)
}
