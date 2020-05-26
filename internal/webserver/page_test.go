package webserver

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var supportedTitles = []string{"contact_us", "our_house", ""}
var expectedResults = map[string]string{
	"contact_us": ": Contact Us",
	"our_house":  ": Our House",
	"":           ""}

func TestPageTitle(t *testing.T) {
	for _, title := range supportedTitles {
		t.Run(fmt.Sprintf("PageTitle=\"%+v\"", title), func(t *testing.T) {
			page := buildPageForTesting(title)
			result := page.PageTitle()

			if result != expectedResults[title] {
				t.Errorf("Expected: \"%+v\" for given title: \"%+v\", got: \"%+v\"", expectedResults[title], title, result)
			}
		})
	}
}

func TestLoadPage(t *testing.T) {
	osWD, _ := os.Getwd()
	workingDirectory = filepath.Clean(filepath.Join(osWD, "\\..\\.."))
	t.Logf("Manually setting working directory to: %+v\n", workingDirectory)

	for _, title := range supportedTitles {
		t.Run(fmt.Sprintf("LoadPage(\"%+v\")", title), func(t *testing.T) {
			page, _ := loadPage(title)
			if page == nil {
				t.Errorf("Couldn't load page with title: %+v", title)
			}
		})
	}
}

func buildPageForTesting(title string) Page {
	return Page{Title: title, Content: "test content"}
}
