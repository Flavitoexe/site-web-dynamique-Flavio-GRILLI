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

	type Product struct {
		Name        string
		Price       float64
		InSale      bool
		PicturePath string
	}
	listProducts := []Product{
		{"Palace Pull A Capuche Unisexe Chasseur", 149.99, false, "/static/img/products/19A.webp"},
		{"Palace Washed Terry 1/4 Placket Hood Mojito", 169.99, false, "/static/img/products/16A.webp"},
		{"Palace Pull A Capuchon Marine", 139.99, false, "/static/img/products/21A.webp"},
		{"Palace Pantalon Bossy Jean Stone", 124.99, false, "/static/img/products/34B.webp"},
		{"Palace Pull Crew Passepose Noir", 129.99, true, "/static/img/products/18A.webp"},
		{"Palace Pantalon Cargo Gore-Tex, R-Tek Noir", 109.99, false, "/static/img/products/33B.webp"},
	}

	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "menu", listProducts)
	})

	path, _ := os.Getwd()
	fileServer := http.FileServer(http.Dir(path + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8000", nil)
}
