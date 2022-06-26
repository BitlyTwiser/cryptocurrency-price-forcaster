package coinGecko

type CoinGecko struct {
	apiKey                string
	coinsListEndpoint     string
	coinsTrendingEndpoint string
	apiVersion            string
	baseUrl               string
}

type TrendingData struct {
	Coins []struct {
		Item struct {
			ID            string  `json:"id"`
			CoinID        int     `json:"coin_id"`
			Name          string  `json:"name"`
			Symbol        string  `json:"symbol"`
			MarketCapRank int     `json:"market_cap_rank"`
			Thumb         string  `json:"thumb"`
			Small         string  `json:"small"`
			Large         string  `json:"large"`
			Slug          string  `json:"slug"`
			PriceBtc      float64 `json:"price_btc"`
			Score         int     `json:"score"`
		} `json:"item"`
	} `json:"coins"`
}

type ListCoinData []struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// type TrendingData struct {
// 	Coins []struct {
// 		Item struct {
// 			ID            string  `json:"id"`
// 			CoinID        int     `json:"coin_id"`
// 			Name          string  `json:"name"`
// 			Symbol        string  `json:"symbol"`
// 			MarketCapRank int     `json:"market_cap_rank"`
// 			Thumb         string  `json:"thumb"`
// 			Small         string  `json:"small"`
// 			Large         string  `json:"large"`
// 			Slug          string  `json:"slug"`
// 			PriceBtc      float64 `json:"price_btc"`
// 			Score         int     `json:"score"`
// 		} `json:"item"`
// 	} `json:"coins"`
// 	Exchanges []interface{} `json:"exchanges"`
// }
