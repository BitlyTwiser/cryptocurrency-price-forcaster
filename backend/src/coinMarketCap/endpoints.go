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
}

func GetCryptoSymbols(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting the crypto data")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map", nil)
	if err != nil {
		log.Print(err)
	}

	q := url.Values{}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("API_KEY"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to server")
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)

	w.Write([]byte(respBody))
}
