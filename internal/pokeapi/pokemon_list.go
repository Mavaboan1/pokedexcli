package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListAreaEncounter(locationName string) (RespAreaLocation, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		encounterResp := RespAreaLocation{}
		err := json.Unmarshal(val, &encounterResp)
		if err != nil {
			return RespAreaLocation{}, err
		}
		return encounterResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaLocation{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespAreaLocation{}, err
	}

	encounterResp := RespAreaLocation{}
	err = json.Unmarshal(dat, &encounterResp)
	if err != nil {
		return RespAreaLocation{}, err
	}

	c.cache.Add(url, dat)
	return encounterResp, nil

}
