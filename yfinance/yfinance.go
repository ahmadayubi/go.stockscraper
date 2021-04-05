package yfinance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://query1.finance.yahoo.com//v10/finance/quoteSummary/%s?modules=earnings,price,calendarEvents,financialData,insiderTransactions,recommendationTrend,majorHoldersBreakdown"

func getEarnings (){

}

func getNews(){

}

func getPriceHistory(){

}

func getFinancialData(){

}

func getInsiderData(){

}

func getRecommendation(){

}

func getData(symbol string) StockResponse{
	req, err := http.NewRequest("GET", fmt.Sprintf(baseURL, symbol), nil)
	if err != nil {
		return StockResponse{}
	}

	client := &http.Client{
		Timeout: time.Second * 20,
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("User-Agent", "Yama_Scraper/1.0")

	var jsonResp YahooResponse
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return StockResponse{}
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(respBody, &jsonResp); err != nil {
		return StockResponse{}
	}

	return StockResponse{}
}

func formatResponse() StockResponse{
	return StockResponse{}
}