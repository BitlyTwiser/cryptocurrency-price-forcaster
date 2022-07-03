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

type probabilityData struct {
	LowProbability     float64 `json:"low_probability"`
	LowPriceEstimate   float64 `json:"low_price_estimate"`
	HighPriceEstimate  float64 `json:"high_price_estimate"`
	HighProbability    float64 `json:"high_probability"`
	CurrentPrice       float64 `json:"current_price"`
	CurrentProbability float64 `json:"current_probability"`
}

//Get the last 30 days of OLCH data from the endpoint and use that for classification
func GetFutureCostPrediction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting started")
	// var returnData = make(map[string]float64)
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

	// val, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", a.naiveGuassianBayesAlgorithm(a.currentTokenPrice)), 64)

	// returnData["probability"] = val
	json.NewEncoder(w).Encode(a.calculateHighMidLowPrices())
}

// Used for future prediction. Uses standard deviation to curate a high,mid, low guestimation.
func (alg *algorithm) calculateHighMidLowPrices() probabilityData {
	deviation := alg.standardDeviation()
	min := 0.0
	max := 0.0

	for _, val := range alg.classificationData {
		if val < min {
			min = val
		}
		if val < max {
			max = val
		}
	}

	valMin, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", alg.naiveGuassianBayesAlgorithm(min-deviation)), 64)
	valMax, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", alg.naiveGuassianBayesAlgorithm(max-deviation)), 64)
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", alg.naiveGuassianBayesAlgorithm(alg.currentTokenPrice)), 64)
	// Note: Need to fix price estimate
	var p = probabilityData{
		LowProbability:     valMin,
		LowPriceEstimate:   min - deviation,
		HighProbability:    valMax,
		HighPriceEstimate:  max + deviation,
		CurrentProbability: val,
		CurrentPrice:       alg.currentTokenPrice,
	}

	return p
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

func (alg *algorithm) naiveGuassianBayesAlgorithm(price float64) float64 {
	mean := alg.mean()
	standardDeviation := alg.standardDeviation()

	return 1 / (math.Sqrt(2*math.Pi) * (standardDeviation)) * math.Pow(math.E, (-math.Pow((price-mean), 2)/(2*(math.Pow(standardDeviation, 2)))))
}

func (alg *algorithm) sumArray(items ...float64) float64 {
	var result float64 = 0
	for _, i := range items {
		result += i
	}

	return result
}
