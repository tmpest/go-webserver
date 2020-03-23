package webserver

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type recaptchaBody struct {
	Token string `json:"token"`
}

type recaptchaValidationResponseBody struct {
	Success       bool     `json:"success"`
	ChallengeTime string   `json:"challenge_ts"`
	Hostname      string   `json:"hostname"`
	Errors        []string `json:"error-codes"`
}

var recaptchaSecret string = os.Getenv("GOOGLE_RECAPTCHA_SECRET")

// recaptchaHandler handles POST request to reCaptcha enpoint, it requires the token result of the reCaptcha request for validation
// if valid, the handler will render a page on success
func recaptchaHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var body recaptchaBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Unable to parse the request!\n%+v\n", err)
		return
	}

	params := url.Values{}
	params.Set("secret", recaptchaSecret)
	params.Set("response", body.Token)
	if err != nil {
		log.Printf("Recieved an error when marshalling our request to Google's Recaptcha Validation API\n%+v\n", err)
		return
	}

	response, err := http.Post("https://www.google.com/recaptcha/api/siteverify?"+params.Encode(), "application/json", nil)
	if err != nil {
		log.Printf("Recieved an error from calling Google's Recaptcha Validation API\n%+v\n", err)
		return
	}

	defer response.Body.Close()
	var responseBody recaptchaValidationResponseBody
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		log.Printf("Unable to parse the response from Google's Validation API call!\n%+v\n", err)
		return
	}

	if !responseBody.Success {
		log.Printf("Response was not a success... User may be a bot!\n%+v\n", responseBody)
		return
	}

	p, err := loadContactUsPage()
	if err != nil {
		return
	}

	defaultTemplate.Execute(w, p)
}

var contactUsPage string = workingDirectory + "/private/contact_us.html"
var contactUsTitle string = "contact_us"

func loadContactUsPage() (*Page, error) {
	content, err := ioutil.ReadFile(contactUsPage)
	if err != nil {
		log.Printf("Recieved an error while reading file: %+v\n%+v\n", contactUsPage, err)
		return nil, err
	}
	return &Page{Title: contactUsTitle, Content: template.HTML(content)}, nil
}
