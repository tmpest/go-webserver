package webserver

import (
	"log"
	"net/http"
)

// Start starts up the webserver
func Start() {
	http.HandleFunc("/", viewHandler)
	http.Handle("/public/assets/images/", http.FileServer(http.Dir("")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
