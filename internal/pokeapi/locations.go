package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) FetchLocations(url *string) (LocationData, error) {
	locationsUrl := APIConfig.BaseURL + APIConfig.APIVersion + "/location-area"
	if url != nil {
		locationsUrl = *url
	}

	req, err := http.NewRequest("GET", locationsUrl, nil)
	if err != nil {
		return LocationData{}, fmt.Errorf("error forming request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, fmt.Errorf("error fetching data: %v", err)
	}
	defer res.Body.Close()

	var locationData LocationData
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationData); err != nil {
		return LocationData{}, fmt.Errorf("failed to parse locations: %v", err)
	}

	return locationData, nil
}
