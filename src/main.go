package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Desc        string
	InSale      bool
	PicturePath string
}

func (prod Product) Apply50Discount() string {
	result := prod.Price / 2
	return fmt.Sprintf("%.2f", result)
}

func main() {

	listTemplates, errTemplate := template.ParseGlob("./templates/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	listProducts := []Product{
		{1, "Palace Pull A Capuche Unisexe Chasseur", 150, "Un hoodie confortable et polyvalent au design sobre. Parfait pour un look streetwear élégant et décontracté.", false, "/static/img/products/19A.webp"},
		{2, "Palace Washed Terry 1/4 Placket Hood Mojito", 170, "Sweat à capuche léger en coton lavé, finition “washed” pour un effet vintage. Teinte mojito rafraîchissante pour l'été.", false, "/static/img/products/16A.webp"},
		{3, "Palace Pull A Capuchon Marine", 140, "Classique intemporel signé Palace, coloris bleu marine profond. Confort maximal pour un style urbain discret.", false, "/static/img/products/21A.webp"},
		{4, "Palace Pantalon Bossy Jean Stone", 125, "Jean coupe droite au ton stone délavé. Idéal pour un look décontracté sans négliger la qualité du denim Palace.", false, "/static/img/products/34B.webp"},
		{5, "Palace Pull Crew Passepose Noir", 130, "Sweat col rond noir avec détails passepoilés contrastants. Allie minimalisme et élégance streetwear.", true, "/static/img/products/18A.webp"},
		{6, "Palace Pantalon Cargo Gore-Tex, R-Tek Noir", 110, "Pantalon technique résistant à l’eau grâce au tissu Gore-Tex. Parfait mélange entre performance et style utilitaire.", false, "/static/img/products/33B.webp"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/" {
			// Gestion d'erreur à faire
		}
		listTemplates.ExecuteTemplate(w, "menu", listProducts)
	})

	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		idStr := strings.TrimPrefix(path, "/product/")
		idInt, err := strconv.Atoi(idStr)

		if err != nil {
			http.NotFound(w, r)
			return
		}
		var selected Product
		for _, p := range listProducts {
			if p.ID == idInt {
				selected = p
				break
			}
		}
		if selected.ID == 0 {
			http.NotFound(w, r)
			return
		}

		listTemplates.ExecuteTemplate(w, "product", selected)
	})

	path, _ := os.Getwd()
	fileServer := http.FileServer(http.Dir(path + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8000", nil)
}
