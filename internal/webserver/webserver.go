package webserver

import (
	"log"
	"net/http"
)

// Start starts up the webserver
func Start() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/contact_us/recaptcha", recaptchaHandler)
	http.Handle("/public/assets/images/", http.FileServer(http.Dir(workingDirectory)))
	http.Handle("/public/javascript/", http.FileServer(http.Dir(workingDirectory)))
	http.Handle("/public/stylesheets/", http.FileServer(http.Dir(workingDirectory)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
