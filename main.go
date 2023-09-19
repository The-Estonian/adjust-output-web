package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"01.kood.tech/git/jsaar/go-reloaded/ascii-art-web/banners"
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
	dataRowArr := compileDataRows(formDataText, formDataBanner)

	dataStruct := struct {
		DataRow  string
		DataRow2 string
	}{
		DataRow:  dataRowArr[0],
		DataRow2: dataRowArr[1],
	}

	tmpl.Execute(w, dataStruct)
}

func compileDataRows(formDataText, formDataBanner string) [2]string {
	if len(formDataText) < 1 {
		return [2]string{}
	}

	// Get carriage return Idx if there is one
	var newLineIdx int
	for i, r := range formDataText {
		if r == 13 {
			newLineIdx = i
			break
		}
	}
	var dataRows [2]string

	// if there is no carriage return, user input is one line (set 2nd line string to be empty)
	if newLineIdx == 0 {
		dataRows[0] = formDataText
		dataRows[1] = ""

		// if there is a carriage return, user input is multi-line
	} else {
		dataRows[0] = formDataText[:newLineIdx]
		dataRows[1] = formDataText[newLineIdx+2:]
	}

	// if user input is one line
	if dataRows[1] == "" {
		dataRows[0] = createAsciiString(dataRows[0], formDataBanner)
		return dataRows

		// if user input is multi-line
	} else {
		dataRows[0] = createAsciiString(dataRows[0], formDataBanner)
		dataRows[1] = createAsciiString(dataRows[1], formDataBanner)
	}

	return dataRows
}

func createAsciiString(text, banner string) string {
	encodedTextArray := banners.EncodeText(text, banner)
	asciiString := ""
	for index, item := range encodedTextArray {
		asciiString += item
		if index < len(encodedTextArray) {
			asciiString += "\n"
		}
	}

	return asciiString
}

func checkNewLine(s string) bool {
	if strings.Contains(s, "\n") {
		return true
	}
	return false
}
