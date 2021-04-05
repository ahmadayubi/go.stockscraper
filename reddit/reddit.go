package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"unicode"
)

const listingURL = "https://www.reddit.com/r/%s/%s.json?limit=%d"

type Response struct {
	Data struct {
		Listings []struct{
			Listing Listing `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type Listing struct {
	Title string `json:"title"`
	Body string `json:"selftext"`
	Ups int `json:"ups"`
	Downs int `json:"downs"`
	Created float64 `json:"created"`
}

type Ticker struct {
	Count int `json:"count"`
	Ups int `json:"ups"`
	Downs int `json:"downs"`
}

type SubReddits struct {
	Name string `json:"name"`
	Sort string `json:"sort"`
	Tickers map[string]Ticker `json:"tickers"`
}

func GetTickers (subreddit, sort string, limit int) SubReddits {
	var resp Response
	if err := listingRequest(subreddit, sort, limit, &resp); err != nil{
		return SubReddits{}
	}
	return SubReddits{
		Name: subreddit,
		Sort: sort,
		Tickers: getTickers(resp),
	}
}

func listingRequest(subreddit, sort string, limit int, jsonResp *Response) error{
	req, err := http.NewRequest("GET", fmt.Sprintf(listingURL, subreddit, sort, limit), nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: time.Second * 20,
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("User-Agent", "Yama_Scraper/1.0")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(respBody, jsonResp); err != nil {
		return err
	}
	return nil
}

func getTickers(listings Response) map[string]Ticker {
	counter := make(map[string]Ticker)
	for _, obj := range listings.Data.Listings{
		l := obj.Listing
		words := strings.Split(l.Title, " ")
		words = append(words, strings.Split(l.Body, " ")...)
		for _, word := range words {
			if isTicker(strings.Trim(word, " ")){
				symbol := strings.ToUpper(strings.Trim(word, " "))
				if _, ok := counter[symbol]; ok {
					ticker := Ticker{
						Count: 1 + counter[symbol].Count,
						Ups: l.Ups + counter[symbol].Ups,
						Downs: l.Downs + counter[symbol].Downs,
					}
					counter[symbol] = ticker
				} else {
					counter[symbol] = Ticker{
						Count: 1,
						Ups: l.Ups,
						Downs: l.Downs,
					}
				}
			}
		}
	}
	for k, v := range counter {
		if v.Count == 1 {
			delete(counter, k)
		}
	}
	return counter
}

func isTicker(s string) bool {
	if len(s) < 2 || len(s) > 4 {
		return false
	}

	isLower := false
	isProbablyTicker := false
	isValid := true
	oddCharCount := 0
	for i, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			isLower = true
		}
		if r == 36 || r == 46 { //$ unicode
			oddCharCount++
			isProbablyTicker = true
		}
		if r == 46 && i >= len(s) - 1 {
			isProbablyTicker = false
			break
		}
		if !unicode.IsLetter(r) && r != 46 && r != 36 {
			isValid = false
			break
		}
	}
	return (isProbablyTicker || !isLower) && isValid && oddCharCount <= 1
}

func SumTotal(sr []SubReddits) SubReddits{
	total := make(map[string]Ticker)
	for _, r := range sr {
		for sym, val := range r.Tickers {
			if _, ok := total[sym]; ok {
				ticker := Ticker{
					Count: val.Count + total[sym].Count,
					Ups: val.Ups + total[sym].Ups,
					Downs: val.Downs + total[sym].Downs,
				}
				total[sym] = ticker
			} else {
				total[sym]= Ticker{
					val.Count,val.Ups,val.Downs,
				}
			}

		}
	}

	for k, v := range total {
		if v.Count < 10 || v.Ups < 20 {
			delete(total, k)
		}
	}
	return SubReddits{
		Name: "total",
		Sort: "sum",
		Tickers: total,
	}

}