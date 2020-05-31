package webserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"text/template"
)

// HousePage data structure for the 'Our House' page
type HousePage struct {
	PageTitle string   `json:"pageTitle"`
	Timeline  Timeline `json:"timeline"`
}

// Timeline data structure for the timeline component
type Timeline struct {
	Entries []TimelineEntry `json:"entries"`
}

// TimelineEntry data structure for the individual timeline entries
type TimelineEntry struct {
	ID    string          `json:"id"`
	Title string          `json:"title"`
	Date  string          `json:"date"`
	Body  []string        `json:"body"`
	Media []TimelineMedia `json:"media"`
}

// TimelineMedia data structure for media elements
type TimelineMedia struct {
	ID       string
	Index    int    `json:"index"`
	Caption  string `json:"caption"`
	AltText  string `json:"alt"`
	Location string `json:"location"`
	Type     string `json:"type"`
}

var ourHouseStateFile = workingDirectory + "/internal/resources/our-house-state.json"

func loadOurHouseData() *HousePage {
	var ourHousePage HousePage
	reader, e := ioutil.ReadFile(ourHouseStateFile)

	if e != nil {
		fmt.Printf("%+v\n", e)
		return nil
	}

	e = json.Unmarshal(reader, &ourHousePage)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return nil
	}

	for _, te := range ourHousePage.Timeline.Entries {
		// Copy the ID from the Parent
		for i, tm := range te.Media {
			tm.ID = te.ID
			te.Media[i] = tm
		}
		// Order the Media Elements according to the indexes
		sort.Slice(te.Media, func(i int, j int) bool {
			return te.Media[i].Index < te.Media[j].Index
		})
	}

	return &ourHousePage
}

// Content returns page contents to be rendered along with the website's template
func (hp HousePage) Content() string {
	return hp.Timeline.ToHTML()
}

var timelineTemplate = workingDirectory + "/private/timeline.html"

// ToHTML returns the HTML for a Timeline element as a string
func (tl Timeline) ToHTML() string {
	t, e := template.New("timeline.html").ParseFiles(timelineTemplate)

	if e != nil {
		fmt.Printf("%+v\n", e)
		return "ERROR!"
	}

	var result bytes.Buffer

	e = t.Execute(&result, tl)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return "ERROR!"
	}

	return result.String()
}

var timelineEntryTemplate = workingDirectory + "/private/timeline_entry.html"

// ToHTML returns the HTML for a TimelineEntry elements as a string
func (te TimelineEntry) ToHTML() string {
	t, e := template.New("timeline_entry.html").ParseFiles(timelineEntryTemplate)

	if e != nil {
		fmt.Printf("%+v\n", e)
		return "ERROR!"
	}

	var result bytes.Buffer

	e = t.Execute(&result, te)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return "ERROR!"
	}

	return result.String()
}

var timelineMediaTemplate = workingDirectory + "/private/timeline_media.html"

// ToHTML returns the HTML for a TimelineMedia elements as a string
func (tm TimelineMedia) ToHTML() string {
	t, e := template.New("timeline_media.html").ParseFiles(timelineMediaTemplate)

	if e != nil {
		fmt.Printf("%+v\n", e)
		return "ERROR!"
	}

	var result bytes.Buffer

	e = t.Execute(&result, tm)
	if e != nil {
		fmt.Printf("%+v\n", e)
		return "ERROR!"
	}

	return result.String()
}

// viewHandler does not used a cached template because it makes development easier
func ourHouseHandler(w http.ResponseWriter, r *http.Request) {
	p := loadOurHouseData()

	template.Must(template.ParseFiles(workingDirectory+"/public/template.html")).Execute(w, p)
}

// IsActivePage accepts a page title and returns true if it matches 'our_house'
func (hp HousePage) IsActivePage(title string) bool {
	if title == "our_house" {
		return true
	}
	return false
}
