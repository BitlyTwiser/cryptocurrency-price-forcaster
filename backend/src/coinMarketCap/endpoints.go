package coinMarketCap

import (
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

	w.Write([]byte(respBody))
}
