package algorithm

import (
	"capstone/src/coinGecko"
	"encoding/json"
	"fmt"
	"math"
	"math/cmplx"
	"net/http"
)

type algorithm struct {
	token                    string
	classificationData       []float64
	eulersConstant           float64
	currentTokenPrice        float64
	classificationDataValues map[string]interface{}
}

type AlgorithmRequestBody struct {
	CoinId string  `json:"coin_id"`
	Price  float64 `json:"price"`
}

//Get the last 30 days of OLCH data from the endpoint and use that for classification
func GetFutureCostPrediction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting started")
	cg := coinGecko.CoinGecko{}
	cg.CoinGeckoConstructor()

	var r AlgorithmRequestBody

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Starting cost prediction analysis for token %s", r.CoinId)
	fmt.Println(math.E)
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(8))
	fmt.Printf("More Complex %v", cmplx.Sqrt(-5*12i))
	fmt.Printf("Exponent: %.1f", math.Pow(6, 2))

	a := algorithm{
		token:              r.CoinId,
		classificationData: cleanClassificationData(cg.GetOHLCDataForToken(r.CoinId)),
		eulersConstant:     math.E,
		currentTokenPrice:  r.Price,
	}

	fmt.Println("Now Printing flat values")
	for _, v := range a.classificationData {
		fmt.Println(v)
	}

	a.train()
	// fmt.Println(a.classificationData)
	// At the end take the first and last price, calculate different, add that to price or subtract depending on if the price is going to move up or down.
}

func loop() {
	items := 10

	for i := 0; i < items; i++ {
		fmt.Println(math.Floor(float64(i / 2)))
	}
}

func cleanClassificationData(classificationData coinGecko.OhlcData) []float64 {
	var flatValues []float64
	for _, day := range classificationData {

		for i, val := range day {
			if i == 0 {
				fmt.Println(fmt.Sprintf("Cleansing data, removing Value: %.1f", day[i]))
			} else {
				flatValues = append(flatValues, val)
			}
		}
	}
	return flatValues
}

func (alg *algorithm) dataClassigicationEngine() {

}

func (alg *algorithm) laplaceSmoothing() {

}

func (alg *algorithm) train() {

}

func (alg *algorithm) probabillityEngine() {

}

func (alg *algorithm) sumArray(items ...float64) float64 {
	var result float64 = 0
	for _, i := range items {
		result += i
	}

	return result
}

func New() {

}
