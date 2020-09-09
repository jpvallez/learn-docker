package apifootball

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//API-FOOTBALL / RAPID-API

type Api struct {
	Wrapper Teams `json:"api"`
}

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	TeamId             int    `json:"team_id"`
	Name               string `json:"name"`
	Code               string `json:"code"`
	Logo               string `json:"logo"`
	Country            string `json:"country"`
	Founded            int    `json:"founded"`
	HomeGround         string `json:"venue_name"`
	HomeGroundCapacity int    `json:"venue_capacity"`
}

func GetTeamsFromLeagueId(leagueId string, api bool) ([]Team, error) {

	var err error
	var body []byte
	var teamsList []Team

	if api == true {
		//TODO: Pay for API-FOOTBALL, get API keys and implement this part properly

		url := "https://api-football-v1.p.rapidapi.com/v2/teams/league/" + leagueId

		req, e := http.NewRequest("GET", url, nil)
		if e != nil {
			//Couldn't make request, return empty team and error
			fmt.Println(e)
			return teamsList, e
		}

		// API keys
		// req.Header.Add("x-rapidapi-host", "")
		// req.Header.Add("x-rapidapi-key", "")

		res, e := http.DefaultClient.Do(req)
		if e != nil {
			//Couldn't get response, return empty team and error
			fmt.Println(e)
			return teamsList, e
		}

		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)

	} else {
		jsonFile, e := os.Open("example-results/teamsbyleagueid.json")
		if e != nil {
			//Couldn't open file on disk, return empty team and error
			fmt.Println(e)
			return teamsList, e
		}
		fmt.Println("Successfully Opened json")

		defer jsonFile.Close()

		body, err = ioutil.ReadAll(jsonFile)
	}

	var apiContent Api

	// Apply JSON content to API handling struct
	json.Unmarshal(body, &apiContent)

	return apiContent.Wrapper.Teams, err

}
