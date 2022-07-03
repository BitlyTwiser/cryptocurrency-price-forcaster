package coinGecko

type CoinGecko struct {
	apiKey                string
	coinsListEndpoint     string
	coinsEndpoint         string
	coinsTrendingEndpoint string
	ohlchEndpoint         string
	apiVersion            string
	baseUrl               string
}

type OhlcData [][]float64

type TrendingData struct {
	Coins []struct {
		Item struct {
			ID            string  `json:"id"`
			CoinID        int     `json:"coin_id"`
			Name          string  `json:"name"`
			Symbol        string  `json:"symbol"`
			MarketCapRank int     `json:"-"`
			Thumb         string  `json:"-"`
			Small         string  `json:"-"`
			Large         string  `json:"-"`
			Slug          string  `json:"-"`
			PriceBtc      float64 `json:"-"`
			Score         int     `json:"score"`
		} `json:"item"`
	} `json:"coins"`
}

type ListCoinData []struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CoinData struct {
	ID                  string  `json:"id"`
	Symbol              string  `json:"symbol"`
	Name                string  `json:"name"`
	PublicInterestScore float64 `json:"public_interest_score"`
	MarketData          struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
		TotalVolume struct {
			Usd float64 `json:"usd"`
		} `json:"total_volume"`
		High24H struct {
			Usd float64 `json:"usd"`
		} `json:"high_24h"`
		Low24H struct {
			Usd float64 `json:"usd"`
		} `json:"low_24h"`
		PriceChange24H    float64 `json:"price_change_24h"`
		TotalSupply       float64 `json:"total_supply"`
		MaxSupply         float64 `json:"max_supply"`
		CirculatingSupply float64 `json:"circulating_supply"`
	} `json:"market_data"`
}
