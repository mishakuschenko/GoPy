package internal

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		log.Fatal("File not found")
	}
	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/login.html")
	if err != nil {
		log.Fatal("File not found")
	}
	temp.Execute(w, nil)
}

func Reg(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../templates/reg.html")
	if err != nil {
		log.Fatal("File not found")
	}
	temp.Execute(w, nil)
}
