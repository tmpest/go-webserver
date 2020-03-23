package webserver

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Page representation for a web page
type Page struct {
	Title   string
	Content template.HTML
}

var pageTitleMap = map[string]string{
	"contact_us":  "Contact Us",
	"our_house":   "Our House",
	"adventures":  "Our Adventures",
	"livys_space": "Livy's Space",
}

// PageTitle returns a formated string to be used when setting the HTML page's title
func (p Page) PageTitle() string {
	if p.Title != "" {
		return fmt.Sprintf(": %+v", pageTitleMap[p.Title])
	}
	return p.Title
}

// IsActivePage accepts a page title and returns true if it matches the current page's title
func (p Page) IsActivePage(title string) bool {
	return p.Title == title
}

// loadPage defaults to loading the public/index.html for empty title values, otherwise loads "public/title_value.html"
func loadPage(title string) (*Page, error) {
	if title == "" {
		title = "index"
	}
	filename := workingDirectory + "/public/" + title + ".html"
	log.Printf("Loading %+v\n", filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Recieved an error while reading file: %+v\n%+v\n", filename, err)
		return nil, err
	}
	return &Page{Title: title, Content: template.HTML(content)}, nil
}

var workingDirectory = getWorkingDirectory()

func getWorkingDirectory() string {
	path := os.Getenv("WEBSERVER_ROOT")
	if path == "" {
		path, _ = os.Getwd()
	}
	return path
}

// defaultTemplate loads the default template.html file for rendering most pages on the website
var defaultTemplate = template.Must(template.ParseFiles(workingDirectory + "/public/template.html"))

// viewHandler simple handler that returns a page based on the url
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, err := loadPage(title)
	if err != nil {
		return
	}

	defaultTemplate.Execute(w, p)
}
