package webserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Page representation for a web page
type Page struct {
	Content []byte
}

// func loadPage(title string) (*Page, error) {
// 	filename := title + ".html"
// 	content, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Page{Content: content}, nil
// }

func loadPage() (*Page, error) {
	filename := "public/index.html"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Content: content}, nil
}

// viewHandler simple handler that returns a page based on the url
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/personal/"):]
	p, _ := loadPage()
	fmt.Fprintf(w, "%s", p.Content)
}
