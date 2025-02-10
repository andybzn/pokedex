package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(url *string) (LocationData, error) {
	locationsUrl := APIConfig.BaseURL + APIConfig.APIVersion + "/location-area"
	if url != nil {
		locationsUrl = *url
	}

	// Hit the pokecache for the url to see if we have this data already
	// If we don't, reach out to the pokeapi to get it
	data, exists := c.cache.Get(locationsUrl)
	if !exists {
		req, err := http.NewRequest("GET", locationsUrl, nil)
		if err != nil {
			return LocationData{}, fmt.Errorf("error forming request: %v", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationData{}, fmt.Errorf("error fetching data: %v", err)
		}
		defer res.Body.Close()

		resData, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationData{}, fmt.Errorf("error reading response: %v", err)
		}

		c.cache.Add(locationsUrl, resData) // cache the response data
		data = resData                     // do the ol' switcharoo
	}

	var locationData LocationData
	if err := json.Unmarshal(data, &locationData); err != nil {
		return LocationData{}, fmt.Errorf("failed to parse locations: %v", err)
	}

	return locationData, nil
}

func (c *Client) ExploreLocation(location string) (LocationArea, error) {
	locationUrl := APIConfig.BaseURL + APIConfig.APIVersion + "/location-area/" + location

	data, exists := c.cache.Get(locationUrl)
	if !exists {
		req, err := http.NewRequest("GET", locationUrl, nil)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error forming request: %v", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error fetching data: %v", err)
		}
		defer res.Body.Close()

		resData, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error reading response: %v", err)
		}

		c.cache.Add(locationUrl, resData)
		data = resData
	}

	var locationAreaData LocationArea
	if err := json.Unmarshal(data, &locationAreaData); err != nil {
		return LocationArea{}, fmt.Errorf("failed to parse pokemon: %v", err)
	}

	return locationAreaData, nil
}
