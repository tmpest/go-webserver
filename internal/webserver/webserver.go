package webserver

import (
	"log"
	"net/http"
)

// Start starts up the webserver
func Start() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
