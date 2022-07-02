package algorithm

import (
	"capstone/src/coinGecko"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
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
	var returnData = make(map[string]float64)
	cg := coinGecko.CoinGecko{}
	cg.CoinGeckoConstructor()

	var r AlgorithmRequestBody

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Starting cost prediction analysis for token %s", r.CoinId)

	a := algorithm{
		token:              r.CoinId,
		classificationData: cleanClassificationData(cg.GetOHLCDataForToken(r.CoinId)),
		eulersConstant:     math.E,
		currentTokenPrice:  r.Price,
	}

	val, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", a.naiveGuassianBayesAlgorithm()), 64)

	returnData["probability"] = val
	json.NewEncoder(w).Encode(returnData)
}

func cleanClassificationData(classificationData coinGecko.OhlcData) []float64 {
	var flatValues []float64
	fmt.Println("Cleansing Data..")
	for _, day := range classificationData {

		for i, val := range day {
			if i == 0 {

			} else {
				flatValues = append(flatValues, val)
			}
		}
	}
	return flatValues
}

func (alg *algorithm) mean() float64 {
	sum := alg.sumArray(alg.classificationData...)

	return math.Ceil((sum / float64(len(alg.classificationData))))
}

func (alg *algorithm) standardDeviation() float64 {
	mean := alg.mean()
	var variance float64

	for _, val := range alg.classificationData {
		variance += math.Pow((val - mean), 2)
	}

	return math.Sqrt(variance / float64(len(alg.classificationData)))
}

func (alg *algorithm) naiveGuassianBayesAlgorithm() float64 {
	mean := alg.mean()
	standardDeviation := alg.standardDeviation()

	return 1 / (math.Sqrt(2*math.Pi) * (standardDeviation)) * math.Pow(math.E, (-math.Pow((alg.currentTokenPrice-mean), 2)/(2*(math.Pow(standardDeviation, 2)))))
}

func (alg *algorithm) sumArray(items ...float64) float64 {
	var result float64 = 0
	for _, i := range items {
		result += i
	}

	return result
}
