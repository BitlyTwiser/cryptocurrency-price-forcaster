package coinGecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (cg *CoinGecko) CoinGeckoConstructor() {
	cg.apiVersion = "v3"
	cg.baseUrl = "https://api.coingecko.com/api"
	cg.coinsEndpoint = "coins"
}

func (cg *CoinGecko) ListCryptoCurrencies(w http.ResponseWriter, req *http.Request) {
	td := TrendingData{}
	fmt.Println(td.Coins)
	log.Println("Getting the crypto data")

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s/list", cg.baseUrl, cg.apiVersion, cg.coinsEndpoint), nil)
	if err != nil {
		log.Print(err)
	}

	q := url.Values{}

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to server")
	}
	fmt.Println(resp.Status)

	respBody, _ := ioutil.ReadAll(resp.Body)

	var data ListCoinData
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		log.Println("Error unmarshalling data. Error: ", err)
	}

	json.NewEncoder(w).Encode(data)
}

func ListTrendinCurrencies() {
	td := TrendingData{}
	fmt.Println(td.Coins)
}
