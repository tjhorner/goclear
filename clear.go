package clear

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

const clearAPIBase string = "https://clear.codeday.org/api/"

// GetCurrentRegions gets all the current Clear regions
func GetCurrentRegions() []Region {
	response, _ := http.Get(clearAPIBase + "regions")
	body, _ := ioutil.ReadAll(response.Body)

	var regions []Region
	json.Unmarshal(body, &regions)

	return regions
}

// GetRegion gets a single specified Clear region
func GetRegion(name string) Region {
	response, _ := http.Get(clearAPIBase + "region/" + template.URLQueryEscaper(name))
	body, _ := ioutil.ReadAll(response.Body)

	var region Region
	json.Unmarshal(body, &region)

	return region
}
