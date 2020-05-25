package webserver

import (
	"log"
	"net/http"
)

// Start starts up the webserver, set useCachedTemplate to true for production
func Start(useCachedTemplate bool) {
	if useCachedTemplate {
		http.HandleFunc("/", cachedTemplateViewHandler)
	} else {
		http.HandleFunc("/", viewHandler)
	}

	http.HandleFunc("/contact_us/recaptcha", recaptchaHandler)
	http.Handle("/public/assets/", http.FileServer(http.Dir(workingDirectory)))
	http.Handle("/public/javascript/", http.FileServer(http.Dir(workingDirectory)))
	http.Handle("/public/stylesheets/", http.FileServer(http.Dir(workingDirectory)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
