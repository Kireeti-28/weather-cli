package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://api.openweathermap.org"

func (c *Client) GetCityWeather(cityName string) (CityResp, error) {
	reqURL := baseURL + "/data/2.5/weather?APPID=" + getApiKey() + "&q=" + cityName

	req, err := http.NewRequest("GET", reqURL, nil)

	if err != nil {
		return CityResp{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return CityResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		type Error struct {
		}
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return CityResp{}, err
	}

	if resp.StatusCode > 399 {
		errorCityResp := ErrorCityResp{}
		json.Unmarshal(respBody, &errorCityResp)
		return CityResp{}, fmt.Errorf("something went wrong..\nmessage: %v\nstatus code: %v", errorCityResp.Message, errorCityResp.Cod)
	}

	cityResp := CityResp{}

	err = json.Unmarshal(respBody, &cityResp)

	if err != nil {
		return CityResp{}, err
	}

	return cityResp, nil
}

func getApiKey() string {
	type tempStorage struct {
		ApiKey string `json:"apiKey"`
	}

	bytes, err := ioutil.ReadFile(".env")

	if err != nil {
		log.Fatalf("unable to read file")
	}

	tempJson := tempStorage{}

	err = json.Unmarshal(bytes, &tempJson)

	if err != nil {
		log.Fatal(err)
	}

	return tempJson.ApiKey
}
