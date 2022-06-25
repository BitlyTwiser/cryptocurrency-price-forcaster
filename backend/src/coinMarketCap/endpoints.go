package coinMarketCap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Endpoints struct {
	apiKey      string
	mapEndpoint string
	apiVersion  string
	baseUrl     string
}
type DataStruct struct {
	Data []struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Rank   int    `json:"rank"`
	} `json:"data"`
}

func (e *Endpoints) EndpointsConstructor() {
	e.apiKey = os.Getenv("API_KEY")
	e.mapEndpoint = "cryptocurrency/map"
	e.apiVersion = "v1"
	e.baseUrl = "pro-api.coinmarketcap.com"
}

func (e *Endpoints) GetCryptoSymbols(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting the crypto data")

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/%s/%s", e.baseUrl, e.apiVersion, e.mapEndpoint), nil)
	if err != nil {
		log.Print(err)
	}

	q := url.Values{}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", e.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to server")
	}
	fmt.Println(resp.Status)

	respBody, _ := ioutil.ReadAll(resp.Body)

	var data DataStruct
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		log.Println("Error unmarshalling data. Error: ", err)
	}

	json.NewEncoder(w).Encode(data)
}
