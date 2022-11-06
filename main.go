package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ttacon/chalk"
)

const APIBaseURL string = "https://api.coindesk.com"

var client *http.Client

type CurrentPriceType struct {
	ChartName string   `json:"chartName"`
	BPI       BPIType  `json:"bpi"`
	Time      TimeType `json:"time"`
}

type USDType struct {
	Rate string `json:"rate"`
}

type TimeType struct {
	Updated string `json:"updated"`
}

type BPIType struct {
	USD USDType `json:"USD"`
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetBitCointRate() {
	url := APIBaseURL + "/v1/bpi/currentprice.json"

	var currentPrice CurrentPriceType

	err := GetJson(url, &currentPrice)

	if err != nil {
		fmt.Printf("error getting BitCoint rate: %s\n", err.Error())
		return
	}

	fmt.Println(chalk.Green.Color("Time:"), currentPrice.Time.Updated)
	fmt.Println(chalk.Green.Color("Rate:"), currentPrice.BPI.USD.Rate, "USD")
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetBitCointRate()
}
