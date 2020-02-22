package webserver

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Page representation for a web page
type Page struct {
	Title   string
	Content template.HTML
}

// PageTitle returns a formated string to be used when setting the HTML page's title
func (p Page) PageTitle() string {
	// fmt.Printf("Recieved Call to: 'PageTitle'\nPage Title: %+v\n", p.Title)
	if p.Title != "" {
		return fmt.Sprintf(": %+v", p.Title)
	}
	return p.Title
}

// IsActivePage accepts a page title and returns true if it matches the current page's title
func (p Page) IsActivePage(title string) bool {
	// fmt.Printf("Recieved Call to: 'IsActivePage' with title: %+v\nPage Title: %+v\n", title, p.Title)
	return p.Title == title
}

// loadPage defaults to loading the public/index.html for empty title values, otherwise loads "public/title_value.html"
func loadPage(title string) (*Page, error) {
	if title == "" {
		title = "index"
	}
	filename := "public/" + title + ".html"
	log.Printf("Loading %+v\n", filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Recieved an error while reading file: %+v\n%+v\n", filename, err)
		return nil, err
	}
	return &Page{Title: title, Content: template.HTML(content)}, nil
}

// viewHandler simple handler that returns a page based on the url
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, err := loadPage(title)
	if err != nil {
		return
	}

	t, err := template.ParseFiles("public/template.html")
	if err != nil {
		log.Printf("Recieved an error while reading template file\n%+v\n", err)
		return
	}

	t.Execute(w, p)
}

// var ImageTemplate string = `<!DOCTYPE html>
// <html lang="en"><head></head>
// <body><img src="data:image/jpg;base64,{{.Image}}"></body>`

// func imageAssetHandler(w http.ResponseWriter, r *http.Request) {
// 	image, err := ioutil.ReadFile(r.URL.Path[len("/"):])
// 	if err != nil {
// 		log.Printf("Recieved an error while reading image file: %+v\n%+v\n", r.URL.Path, err)
// 		return
// 	}
// 	fmt.Fprintf(w, "%+v", image)
// }
