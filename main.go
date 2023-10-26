package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Pairs    []string          `json:"pairs"`
	Info     string            `json:"info"`
	Filtered map[string][]string `json:"filtered"`
	Refresh int `json:"refresh_period"`
}

func getPairs(w http.ResponseWriter, r *http.Request) {
	url := "https://remotepairlist.com/?r=1&filter=noprefilter&sort=exchange_24h_change&exchange=binance&market=futures&stake=USDT&limit=300&exchange=binance&stake_currency=USDT&show=1"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var data Response
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	// 打印交易对
	for _, pair := range data.Pairs {
		fmt.Println(pair)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/pairs", getPairs)
	log.Fatal(http.ListenAndServe(":8000", nil))
}