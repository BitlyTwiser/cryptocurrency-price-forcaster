package main

import (
	"capstone/src/algorithm"
	"capstone/src/coinGecko"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Exported struct to eventually assist with testing of the API.
type App struct {
	router *mux.Router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading environment variables. Error: %v", err)
	}
	run()
}

func run() {
	//Load new router
	mux := newRouter()
	port := os.Getenv("PORT")
	fmt.Println("Listening on ", port)
	serv := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
		Handler: handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{
				"Content-Type", "X-Requested-With", "Access-Control-Allow-Origin", "Origin",
				"Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization",
			}),
		)(mux.router),
	}

	log.Fatal(serv.ListenAndServe())
}

func (a App) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func newRouter() App {
	cg := coinGecko.CoinGecko{}
	cg.CoinGeckoConstructor()

	a := App{router: mux.NewRouter()}
	a.router.Use(mux.CORSMethodMiddleware(a.router))
	a.router.Handle("/favicon.ico", http.NotFoundHandler())
	a.router.HandleFunc("/get-crypto-symbols", cg.ListCryptoCurrencies).Methods(http.MethodGet)
	a.router.HandleFunc("/get-trending-data", cg.ListTrendinCurrencies).Methods(http.MethodGet)
	a.router.HandleFunc("/get-coin-data", cg.GetCoinData).Methods(http.MethodGet)
	a.router.HandleFunc("/get-prediction", algorithm.GetFutureCostPrediction).Methods(http.MethodPost)
	return a
}
