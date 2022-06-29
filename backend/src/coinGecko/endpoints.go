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
	cg.coinsListEndpoint = "coins/list"
	cg.coinsTrendingEndpoint = "search/trending"
	cg.coinsEndpoint = "coins"
}

func (cg *CoinGecko) requestMaker(structToMarshall interface{}, endpoint string) interface{} {
	log.Println("Getting the crypto data")

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", cg.baseUrl, cg.apiVersion, endpoint), nil)
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

	err = json.Unmarshal(respBody, &structToMarshall)
	if err != nil {
		log.Println("Error unmarshalling data. Error: ", err)
	}

	return structToMarshall
}

func (cg *CoinGecko) GetCoinData(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["coin_id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	log.Printf("Getting the data for %s", keys[0])

	var data CoinData
	filledData := cg.requestMaker(data, fmt.Sprintf("%s/%s", cg.coinsEndpoint, keys[0]))

	json.NewEncoder(w).Encode(filledData)
}

func (cg *CoinGecko) ListCryptoCurrencies(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting the crypto data")

	var data ListCoinData
	filledData := cg.requestMaker(data, cg.coinsListEndpoint)

	json.NewEncoder(w).Encode(filledData)
}

func (cg *CoinGecko) ListTrendinCurrencies(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting trending data")

	var data TrendingData
	filledData := cg.requestMaker(data, cg.coinsTrendingEndpoint)

	json.NewEncoder(w).Encode(filledData)
}
