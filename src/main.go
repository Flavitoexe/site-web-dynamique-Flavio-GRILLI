package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {

	listTemplates, errTemplate := template.ParseGlob("./templates/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "/")
	})

	http.ListenAndServe("localhost:8000", nil)

}
