package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func PingCMC() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// карта параметров клюс - массив значений. Добавка параметров к урл
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "50")
	q.Add("convert", "USD")
	// start=1&limit=5000&convert=USD

	// установка загаловков в хедер
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "5bb1d0b7-5ab1-4420-8fb6-1b5b0fad1a4c")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
}
