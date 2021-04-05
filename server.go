package main

import (
	"bytes"
	"encoding/json"
	"go.stockscraper/reddit"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var subreddits = []string{"wallstreetbets", "pennystocks", "CanadianInvestor", "StockMarket", "Daytrading", "stocks"}
var sorts = []string{"hot", "new", "top"}

func main(){
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/json", getJson)
	log.Println("Running...")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func getJson (w http.ResponseWriter, r *http.Request){
	var subArray []reddit.SubReddits
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	subChannel := make(chan reddit.SubReddits)
	defer close(subChannel)

	for _, subreddit := range subreddits {
		for _, sort := range sorts {
			go func(r, s string, l int) {
				subChannel <- reddit.GetTickers(r, s, l)
			}(subreddit, sort, limit)
		}
	}

	for i:= 0; i < len(subreddits) * len(sorts); i++ {
		subArray = append(subArray, <-subChannel)
	}

	subArray = append(subArray, reddit.SumTotal(subArray))

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err = enc.Encode(subArray); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveIndex (w http.ResponseWriter, r *http.Request){
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}