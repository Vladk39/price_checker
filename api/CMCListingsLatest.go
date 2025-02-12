package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Listings struct {
	Data   []Cryptocurrency `json:"data"`
	Status Status           `json:"status"`
}

type Cryptocurrency struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	CmcRank           int       `json:"cmc_rank"`
	MarketPairs       int       `json:"num_market_pairs"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_suply"`
	MCTotalSupply     float64   `json:"market_cap_by_total_supply"`
	Max_supply        float64   `json:"max_supply"`
	InfiniteSupply    bool      `json:"infinite_supply"`
	LastUpdate        time.Time `json:"last_updated"`
	DateAdd           time.Time `json:"date_added"`
	Tags              []string  `json:"-"`
	SelfRepCircSupply float64   `json:"self_reported_circulating_supply"`
	SelfRepMarketCap  float64   `json:"self_reported_market_cap"`
	TvlRatio          float64   `json:"tvl_ratio"`
	// omitempty указатель, так как структура может быть нул
	Platform Platform             `json:"platform,omitempty"`
	Quote    map[string]QuoteData `json:"quote"`
}

type Platform struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Slug        string `json:"slug"`
	TokenAdress string `json:"token_adress"`
}
type QuoteData struct {
	Price                 float64   `json:"price"`
	Volume24h             float64   `json:"volume_24h"`
	VolumeChange24h       float64   `json:"volume_change_24h"`
	Volume24hReported     *float64  `json:"volume_24h_reported,omitempty"` // Может отсутствовать
	Volume7d              *float64  `json:"volume_7d,omitempty"`
	Volume7dReported      *float64  `json:"volume_7d_reported,omitempty"`
	Volume30d             *float64  `json:"volume_30d,omitempty"`
	Volume30dReported     *float64  `json:"volume_30d_reported,omitempty"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	TVL                   *float64  `json:"tvl,omitempty"` // Может быть null
	PercentChange1h       float64   `json:"percent_change_1h"`
	PercentChange24h      float64   `json:"percent_change_24h"`
	PercentChange7d       float64   `json:"percent_change_7d"`
	LastUpdated           time.Time `json:"last_updated"`
}

type Status struct {
	Timestamp   time.Time `json:"timestamp"`
	ErrorCode   int       `json:"error_code"`
	ErrorMsg    string    `json:"error_message"`
	CreditCount int       `json:"credit_count"`
	Notice      string    `json:"string"`
}

type CryptoForDB struct {
}

func ListingsLatest() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// карта параметров клюс - массив значений. Добавка параметров к урл
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "2")
	q.Add("convert", "USD")
	// start=1&limit=5000&convert=USD

	// установка загаловков в хедер
	req.Header.Set("Accepts", "application/json")
	//битое апи
	req.Header.Add("X-CMC_PRO_API_KEY", "5bb1d0b7-5ab1-4420-8fb6-1b5b0fad1a4c")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request")
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)

	// respBody, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(respBody))
	var Listings Listings
	err = json.NewDecoder(resp.Body).Decode(&Listings)
	if err != nil {
		log.Print("Ошибка декодирования ответа,", err)
	}
	// Jsonlist, _ := json.MarshalIndent(Listings, "", "  ")
	// fmt.Println(string(Jsonlist))

	for _, crypto := range Listings.Data {
		fmt.Println("Криптовалюта:", crypto.Name, "(", crypto.Symbol, ")")

		for currency, quote := range crypto.Quote {
			fmt.Printf("Курс в %s: %.2f USD\n", currency, quote.Price)
			fmt.Printf("Изминение цены за 24 часа: %.2f%%\n\n", quote.PercentChange24h)
		}
	}

}
