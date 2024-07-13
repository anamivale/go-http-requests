package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	tmplt, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmplt.Execute(w, nil)

}

func HandleContact(w http.ResponseWriter, r *http.Request) {

	tmplt, err := template.ParseFiles("templates/contact.html")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmplt.Execute(w, nil)

}

func main() {

	http.HandleFunc("/", Handler)
	http.HandleFunc("/contact", HandleContact)
	http.ListenAndServe(":5000", nil)
}