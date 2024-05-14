package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageInformation struct {
	Title       string
	Description string
	Font1       string
	Font2       string
	Font3       string
	Output      string
}

func DefineRoutes() {
	// http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))

	// Define your routes here
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/": func(w http.ResponseWriter, r *http.Request) { HomeHandler(w, r, "views/home.html") },
		"/categores": func(w http.ResponseWriter, r *http.Request) { HomeHandler(w, r, "views/home.html") },
	}
	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request, view string) {
	tmpl, err := template.ParseFiles(view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pageInfo := pageInformation{
		Title:       "Home Page",
		Description: "This is the home page",
		Font1:       "Arial",
		Font2:       "Verdana",
		Font3:       "Times New Roman",
		Output:      "hello",
	}
	err = tmpl.Execute(w, pageInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	if r.Method == "POST" {
		inputValue := r.FormValue("input")
		log.Println("Field value: ", inputValue)
		selectValue := r.FormValue("select")
		log.Println("Select value: ", selectValue)
	}
}

func main() {
	DefineRoutes()
	log.Println("Server is starting...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
