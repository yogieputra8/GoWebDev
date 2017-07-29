package main

import (
	"html/template"
	"net/http"
	"log"
	_ "fmt"
)

var tpl *template.Template

type pageData struct {
	Title string
	FirstName string
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abot)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/apply", aply)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)

	if err != nil {
		log.Println("Logged: ", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	// fmt.Println(req.URL.Path)
	// fmt.Println("We got here")
}

func abot(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func cntct(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func aply(w http.ResponseWriter, req *http.Request){

	pd := pageData{
		Title: "Apply Page",
	}

	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", pd)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}