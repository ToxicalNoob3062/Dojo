package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})
	fmt.Println("Starting front end service on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panic(err)
	}

}

func render(w http.ResponseWriter, t string) {
	templatesDir := "/app/templates"

	partials := []string{
		filepath.Join(templatesDir, "base.layout.gohtml"),
		filepath.Join(templatesDir, "header.partial.gohtml"),
		filepath.Join(templatesDir, "footer.partial.gohtml"),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, filepath.Join(templatesDir, t))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
