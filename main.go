package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	level, err := log.ParseLevel(getProperties().LogLevel)
	if err == nil {
		log.SetLevel(level)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", getProperties().Server.Port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	tickers, err3 := doTheJob()
	if err3 != nil {
		log.Error(err3)
		w.WriteHeader(500)
		_, _ = fmt.Fprint(w, err3.Error())
		return
	}

	result, err1 := json.Marshal(tickers)
	if err1 != nil {
		log.Error(err1)
		w.WriteHeader(500)
		_, _ = fmt.Fprint(w, err1.Error())
		return
	}

	_, err2 := w.Write(result)
	if err2 != nil {
		log.Error(err2)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func doTheJob() (*tickerCollection, error) {
	tickers, errorz := parseOnline()
	if len(errorz) != 0 {
		log.Error(errorz)
		return nil, fmt.Errorf("cannot parse pages, check the logs:\n%s", errorz)
	}

	filteredTickers := filter(tickers)

	return &tickerCollection{Tickers: filteredTickers}, nil
}

func filter(tickers *[]stockTicker) *[]stockTicker {
	var filters []func(*stockTicker) error
	filters = append(filters, filterOldForecasts)
	filters = append(filters, filterExtremeForecasts)

	var filteredTickers []stockTicker
	for _, ticker := range *tickers {
		ok := true
		for _, filter := range filters {
			if err := filter(&ticker); err != nil {
				ok = false
				break
			}
		}
		if ok {
			filteredTickers = append(filteredTickers, ticker)
		}
	}

	for i, ticker := range filteredTickers {
		sum := 0.0
		for _, forecast := range *ticker.Forecasts {
			sum += forecast.ExpectedDiff
		}
		filteredTickers[i].Consensus = sum / float64(len(*ticker.Forecasts))
		log.Infof("%v", ticker)
	}

	return &filteredTickers
}
