package algorithm

import (
	"capstone/src/coinGecko"
	"encoding/json"
	"fmt"
	"math"
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
	// Python algo example.
	// probas.append((1/np.sqrt(2*np.pi*self.stds_[j]**2)*np.exp(-0.5*((X[i]-self.means_[j])/self.stds_[j])**2)).prod()*self.priors_[j])
	// Get mean of dataset and standard deviation for equation.

	fmt.Printf("Starting cost prediction analysis for token %s", r.CoinId)
	// fmt.Println(math.E)
	// fmt.Println(math.Pi)
	// fmt.Println(math.Sqrt(8))
	// fmt.Printf("More Complex %v", cmplx.Sqrt(-5*12i))

	a := algorithm{
		token:              r.CoinId,
		classificationData: cleanClassificationData(cg.GetOHLCDataForToken(r.CoinId)),
		eulersConstant:     math.E,
		currentTokenPrice:  r.Price,
	}

	// fmt.Println("Now Printing flat values")
	// for _, v := range a.classificationData {
	// 	fmt.Println(v)
	// }
	fmt.Printf("Exponent: %.1f\n", math.Pow(6, 2))
	fmt.Printf("Mean: %v\n", a.mean())
	fmt.Printf("Standard Deviation: %v\n", a.standardDeviation())
	fmt.Printf("Probability: %v", a.naiveGuassianBayesAlgorithm())
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

func (alg *algorithm) mean() float64 {
	sum := alg.sumArray(alg.classificationData...)

	return math.Ceil((sum / float64(len(alg.classificationData))))
}

func (alg *algorithm) standardDeviation() float64 {
	// Gather mean
	mean := alg.mean()

	// setting (val - mean) ^ 2 to all elements of array
	for i, val := range alg.classificationData {
		alg.classificationData[i] = math.Pow((val - mean), 2)
	}

	varaiance := alg.sumArray(alg.classificationData...)

	return math.Sqrt(varaiance)
}

// P(x_i \mid y) = \frac{1}{\sqrt{2\pi\sigma^2_y}} \exp\left(-\frac{(x_i - \mu_y)^2}{2\sigma^2_y}\right)
// I think we loop here, build probabbilitis, used trained data to determine the probvabilty of our current price ,then return value
func (alg *algorithm) naiveGuassianBayesAlgorithm() float64 {
	initial := (1 / (math.Sqrt(2 * math.Pi * math.Pow(alg.standardDeviation(), 2))))
	second := (math.Pow(alg.eulersConstant, math.Pow(-0.5*(alg.currentTokenPrice-(alg.mean()/alg.standardDeviation())), 2)))

	probability := initial * second

	return probability
}

func (alg *algorithm) priceCalculation() {

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
