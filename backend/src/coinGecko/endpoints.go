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
	cg.ohlchEndpoint = "ohlc?vs_currency=usd&days=30"
	cg.coinsEndpoint = "coins"
}

func requestMaker[V interface{} | [][]float64](structToMarshall V, endpoint string, cg *CoinGecko, q url.Values) V {
	log.Println("Getting the crypto data")

	client := &http.Client{}
	fmt.Println(fmt.Sprintf("%s/%s/%s", cg.baseUrl, cg.apiVersion, endpoint))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", cg.baseUrl, cg.apiVersion, endpoint), nil)
	if err != nil {
		log.Print(err)
	}

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
		w.Write([]byte("Url Param 'coin_id' is missing"))
		return
	}

	log.Printf("Getting the data for %s", keys[0])

	var data CoinData
	filledData := requestMaker(data, fmt.Sprintf("%s/%s", cg.coinsEndpoint, keys[0]), cg, url.Values{})

	json.NewEncoder(w).Encode(filledData)
}

func (cg *CoinGecko) ListCryptoCurrencies(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting the crypto data")

	var data ListCoinData
	filledData := requestMaker(data, cg.coinsListEndpoint, cg, url.Values{})

	json.NewEncoder(w).Encode(filledData)
}

func (cg *CoinGecko) ListTrendinCurrencies(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting trending data")

	var data TrendingData
	filledData := requestMaker(data, cg.coinsTrendingEndpoint, cg, url.Values{})

	json.NewEncoder(w).Encode(filledData)
}

func (cg *CoinGecko) GetOHLCDataForToken(token string) OhlcData {
	log.Println("Getting OHLC Data")

	q := url.Values{}
	q.Add("days", "30")
	q.Add("vs_currency", "usd")
	var data OhlcData
	filledData := requestMaker(data, fmt.Sprintf("%s/%s/%s", cg.coinsEndpoint, token, cg.ohlchEndpoint), cg, q)

	return filledData
}
