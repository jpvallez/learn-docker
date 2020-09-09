package scrape

import (
	"io/ioutil"
	"net/http"
)

// The idea here was to scrape a website with soccer player statistics
// It scrapes an entire website and returns it as a string.
// We could then process that string to find specific details on players

func scrapeSite(url string) (string, error) {
	// Make request to the url provided
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(dataInBytes), nil

	//indexOfPlayerStats = getIndexInHTMLByRefString("DataStore.prime('ws-stage-stat', ")
	//overallPlayStatsContent := pageContent[overallPlayerStatsIndex:titleEndIndex]

}

// Find index of end of reference string in page content.
// E.g. give me the start of the player stats table which starts with "<table id=\"players\">"
/*
func getIndexInHTMLByRefString(ref string) int {
	//find index start
	index := strings.Index(PageContent, ref)
	if index == -1 {
		os.Exit(0)
	}

	//index starts at beginning of reference string
	index = index + len(ref)
	return index
}
*/
