package main

import (
	"01.kood.tech/git/jsaar/go-reloaded/ascii-art-web/banners"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	staticFiles := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", staticFiles))

	http.HandleFunc("/", ServeTemplate)

	fmt.Println("Running server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		fmt.Println(err)
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	formDataBanner := r.FormValue("banner")
	formDataText := r.FormValue("userInput")
	encodedTextArray := banners.EncodeText(formDataText, formDataBanner)
	returnString := ""

	for index, item := range encodedTextArray {
		returnString += item
		if index < len(encodedTextArray) {
			returnString += "\n"
		}
	}

	dataStruct := struct {
		DataRow string
	}{
		DataRow: returnString,
	}

	tmpl.Execute(w, dataStruct)
}
