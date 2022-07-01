package algorithm

import (
	"capstone/src/coinGecko"
	"fmt"
	"net/http"
)

type algorithm struct {
	token              string
	classificationData [][]float64
}

//Get the last 30 days of OLCH data from the endpoint and use that for classification
func GetFutureCostPrediction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting started")
	cg := coinGecko.CoinGecko{}
	cg.CoinGeckoConstructor()

	keys, ok := req.URL.Query()["coin_id"]

	if !ok || len(keys[0]) < 1 {
		w.Write([]byte("Url Param 'coin_id' is missing"))
		return
	}
	fmt.Printf("Starting cost prediction analysis for token %s", keys[0])

	a := algorithm{token: keys[0]}
	fmt.Println(a)

}

func (alg *algorithm) obtainClassificationData() {

}
