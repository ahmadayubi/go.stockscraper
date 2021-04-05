package yfinance

type YahooResponse struct {
	QuoteSummary struct {
		Result []struct {
			RecommendationTrend struct {
				Trend Recommendation`json:"trend"`
				MaxAge int `json:"maxAge"`
			} `json:"recommendationTrend"`
			MajorHoldersBreakdown struct {
				MaxAge              int `json:"maxAge"`
				InsidersPercentHeld struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"insidersPercentHeld"`
				InstitutionsPercentHeld struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"institutionsPercentHeld"`
				InstitutionsFloatPercentHeld struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"institutionsFloatPercentHeld"`
				InstitutionsCount struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"institutionsCount"`
			} `json:"majorHoldersBreakdown"`
			Earnings struct {
				MaxAge        int `json:"maxAge"`
				EarningsChart struct {
					Quarterly []Quarter `json:"quarterly"`
					CurrentQuarterEstimate struct {
						Raw float64 `json:"raw"`
						Fmt string  `json:"fmt"`
					} `json:"currentQuarterEstimate"`
					CurrentQuarterEstimateDate string `json:"currentQuarterEstimateDate"`
					CurrentQuarterEstimateYear int    `json:"currentQuarterEstimateYear"`
					EarningsDate               []struct {
						Raw int    `json:"raw"`
						Fmt string `json:"fmt"`
					} `json:"earningsDate"`
				} `json:"earningsChart"`
				FinancialsChart struct {
					Yearly []struct {
						Date    int `json:"date"`
						Revenue struct {
							Raw     int64  `json:"raw"`
							Fmt     string `json:"fmt"`
							LongFmt string `json:"longFmt"`
						} `json:"revenue"`
						Earnings struct {
							Raw     int64  `json:"raw"`
							Fmt     string `json:"fmt"`
							LongFmt string `json:"longFmt"`
						} `json:"earnings"`
					} `json:"yearly"`
					Quarterly []struct {
						Date    string `json:"date"`
						Revenue struct {
							Raw     int64  `json:"raw"`
							Fmt     string `json:"fmt"`
							LongFmt string `json:"longFmt"`
						} `json:"revenue"`
						Earnings struct {
							Raw     int64  `json:"raw"`
							Fmt     string `json:"fmt"`
							LongFmt string `json:"longFmt"`
						} `json:"earnings"`
					} `json:"quarterly"`
				} `json:"financialsChart"`
				FinancialCurrency string `json:"financialCurrency"`
			} `json:"earnings"`
			CalendarEvents struct {
				MaxAge   int `json:"maxAge"`
				Earnings struct {
					EarningsDate []struct {
						Raw int    `json:"raw"`
						Fmt string `json:"fmt"`
					} `json:"earningsDate"`
					EarningsAverage struct {
						Raw float64 `json:"raw"`
						Fmt string  `json:"fmt"`
					} `json:"earningsAverage"`
					EarningsLow struct {
						Raw float64 `json:"raw"`
						Fmt string  `json:"fmt"`
					} `json:"earningsLow"`
					EarningsHigh struct {
						Raw float64 `json:"raw"`
						Fmt string  `json:"fmt"`
					} `json:"earningsHigh"`
					RevenueAverage struct {
						Raw     int64  `json:"raw"`
						Fmt     string `json:"fmt"`
						LongFmt string `json:"longFmt"`
					} `json:"revenueAverage"`
					RevenueLow struct {
						Raw     int64  `json:"raw"`
						Fmt     string `json:"fmt"`
						LongFmt string `json:"longFmt"`
					} `json:"revenueLow"`
					RevenueHigh struct {
						Raw     int64  `json:"raw"`
						Fmt     string `json:"fmt"`
						LongFmt string `json:"longFmt"`
					} `json:"revenueHigh"`
				} `json:"earnings"`
				ExDividendDate struct {
					Raw int    `json:"raw"`
					Fmt string `json:"fmt"`
				} `json:"exDividendDate"`
				DividendDate struct {
					Raw int    `json:"raw"`
					Fmt string `json:"fmt"`
				} `json:"dividendDate"`
			} `json:"calendarEvents"`
			Price struct {
				MaxAge          int `json:"maxAge"`
				PreMarketChange struct {
				} `json:"preMarketChange"`
				PreMarketPrice struct {
				} `json:"preMarketPrice"`
				PreMarketSource         string `json:"preMarketSource"`
				PostMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketChangePercent"`
				PostMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketChange"`
				PostMarketTime  int `json:"postMarketTime"`
				PostMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketPrice"`
				PostMarketSource           string `json:"postMarketSource"`
				RegularMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChangePercent"`
				RegularMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChange"`
				RegularMarketTime int `json:"regularMarketTime"`
				PriceHint         struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"priceHint"`
				RegularMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPrice"`
				RegularMarketDayHigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayHigh"`
				RegularMarketDayLow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayLow"`
				RegularMarketVolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"regularMarketVolume"`
				AverageDailyVolume10Day struct {
				} `json:"averageDailyVolume10Day"`
				AverageDailyVolume3Month struct {
				} `json:"averageDailyVolume3Month"`
				RegularMarketPreviousClose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPreviousClose"`
				RegularMarketSource string `json:"regularMarketSource"`
				RegularMarketOpen   struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketOpen"`
				StrikePrice struct {
				} `json:"strikePrice"`
				OpenInterest struct {
				} `json:"openInterest"`
				Exchange              string      `json:"exchange"`
				ExchangeName          string      `json:"exchangeName"`
				ExchangeDataDelayedBy int         `json:"exchangeDataDelayedBy"`
				MarketState           string      `json:"marketState"`
				QuoteType             string      `json:"quoteType"`
				Symbol                string      `json:"symbol"`
				UnderlyingSymbol      interface{} `json:"underlyingSymbol"`
				ShortName             string      `json:"shortName"`
				LongName              string      `json:"longName"`
				Currency              string      `json:"currency"`
				QuoteSourceName       string      `json:"quoteSourceName"`
				CurrencySymbol        string      `json:"currencySymbol"`
				FromCurrency          interface{} `json:"fromCurrency"`
				ToCurrency            interface{} `json:"toCurrency"`
				LastMarket            interface{} `json:"lastMarket"`
				Volume24Hr            struct {
				} `json:"volume24Hr"`
				VolumeAllCurrencies struct {
				} `json:"volumeAllCurrencies"`
				CirculatingSupply struct {
				} `json:"circulatingSupply"`
				MarketCap struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"marketCap"`
			} `json:"price"`
			InsiderTransactions struct {
				Transactions []struct {
					MaxAge int `json:"maxAge"`
					Shares struct {
						Raw     int    `json:"raw"`
						Fmt     string `json:"fmt"`
						LongFmt string `json:"longFmt"`
					} `json:"shares"`
					Value struct {
						Raw     int    `json:"raw"`
						Fmt     string `json:"fmt"`
						LongFmt string `json:"longFmt"`
					} `json:"value,omitempty"`
					FilerURL        string `json:"filerUrl"`
					TransactionText string `json:"transactionText"`
					FilerName       string `json:"filerName"`
					FilerRelation   string `json:"filerRelation"`
					MoneyText       string `json:"moneyText"`
					StartDate       struct {
						Raw int    `json:"raw"`
						Fmt string `json:"fmt"`
					} `json:"startDate"`
					Ownership string `json:"ownership"`
				} `json:"transactions"`
				MaxAge int `json:"maxAge"`
			} `json:"insiderTransactions"`
			FinancialData struct {
				MaxAge       int `json:"maxAge"`
				CurrentPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"currentPrice"`
				TargetHighPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"targetHighPrice"`
				TargetLowPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"targetLowPrice"`
				TargetMeanPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"targetMeanPrice"`
				TargetMedianPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"targetMedianPrice"`
				RecommendationMean struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"recommendationMean"`
				RecommendationKey       string `json:"recommendationKey"`
				NumberOfAnalystOpinions struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"numberOfAnalystOpinions"`
				TotalCash struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"totalCash"`
				TotalCashPerShare struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"totalCashPerShare"`
				Ebitda struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"ebitda"`
				TotalDebt struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"totalDebt"`
				QuickRatio struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"quickRatio"`
				CurrentRatio struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"currentRatio"`
				TotalRevenue struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"totalRevenue"`
				DebtToEquity struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"debtToEquity"`
				RevenuePerShare struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"revenuePerShare"`
				ReturnOnAssets struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"returnOnAssets"`
				ReturnOnEquity struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"returnOnEquity"`
				GrossProfits struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"grossProfits"`
				FreeCashflow struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"freeCashflow"`
				OperatingCashflow struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"operatingCashflow"`
				EarningsGrowth struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"earningsGrowth"`
				RevenueGrowth struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"revenueGrowth"`
				GrossMargins struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"grossMargins"`
				EbitdaMargins struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"ebitdaMargins"`
				OperatingMargins struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"operatingMargins"`
				ProfitMargins struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"profitMargins"`
				FinancialCurrency string `json:"financialCurrency"`
			} `json:"financialData"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteSummary"`
}

type StockResponse struct {
	Recommendation Recommendation `json:"recommendation"`
	Earnings []struct{
		Quarterly []Quarter `json:"quarterly"`
		Current Quarter `json:"current"`
	} `json:"earnings"`
	Holders struct{
		Insider string `json:"insider"`
		Institution string `json:"institution"`
		InstFloat string `json:"institutionFloat"`
		InstCount string `json:"institutionCount"`
	} `json:"holders"`

}

type Quarter struct {
	Date string `json:"date"`
	Actual struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"actual"`
	Estimate struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"estimate"`
}

type Recommendation struct {
	Trend []struct {
		Period     string `json:"period"`
		StrongBuy  int    `json:"strongBuy"`
		Buy        int    `json:"buy"`
		Hold       int    `json:"hold"`
		Sell       int    `json:"sell"`
		StrongSell int    `json:"strongSell"`
	} `json:"trend"`
}