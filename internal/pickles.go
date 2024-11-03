package internal;

import (
	"html/template"
	"net/http"
	"log"
);

func Index(w http.ResponseWriter, r *http.Request) {
	type TitleOfPage struct {
		Title string
	};
	Index := TitleOfPage{
		Title: "GoPy!",
	};
	temp, err := template.ParseFiles("../templates/index.html");
	if err != nil {
		log.Fatal("File not found");
	};
	
	temp.Execute(w, Index);
};


