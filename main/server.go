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

const PORT = "3000"

func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templ, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		log.Fatal(w, "Internal Server Error 500", http.StatusInternalServerError)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return

	}

	err = templ.Execute(w, Data)
	if err != nil {
		log.Fatal(err)
	}
	Data.Text, Data.Content, Data.Warning = "", "", ""
}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return

	}
	if !ascii.Isbanner(r.FormValue("select")) {
		http.Error(w, "Bad Request 400", http.StatusBadRequest)
		return
	}

	Data.Text = r.FormValue("text")
	if !ascii.IsPrintable(Data.Text) {
		Data.Warning = "The text contains non-printable characters."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get the value of a form field and assign it to the Data struct

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
	log.Println("\033[32mServer is running on port " + PORT + "...🚀\033[0m")
	log.Println("\033[32mhttp://localhost:" + PORT + "\033[0m")
	http.ListenAndServe(":"+PORT, nil)
}
