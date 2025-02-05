package api

import (
	"fmt"
	"io"
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
	q.Add("limit", "1")
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
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
