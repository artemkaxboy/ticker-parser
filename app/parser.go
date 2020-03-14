package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const expectedForecastsCount = 5

var (
	priceRegex = regexp.MustCompile(`[^\d,]`)
)

func getPriceValue(rawText string) (float64, error) {
	clearString := priceRegex.ReplaceAllString(rawText, "")
	pointedString := strings.ReplaceAll(clearString, ",", ".")
	value, err := strconv.ParseFloat(pointedString, 64)
	return value, err
}

// Parse extracts satellite items from given reader and sends them to given chData channel.
// Occurred errors are sent to chErr channel. url string is used for tracing purposes only.
func Parse(url string, reader io.Reader, chData chan stockTicker, chErr chan error) {
	log.Debugf("parsing started: %s", url)

	document, err1 := goquery.NewDocumentFromReader(reader)
	if err1 != nil {
		err1 = fmt.Errorf("error reading HTTP response body (%s): %w", url, err1)
		log.Debug(err1)
		chErr <- err1
		return
	}

	ticker := stockTicker{}

	fullNameRaw := document.Find(".header__tool__name-full").Text()
	ticker.Name.Full = strings.TrimSpace(fullNameRaw)

	shortNameRaw := document.Find(".header__tool__name-short").Text()
	ticker.Name.Short = strings.TrimSpace(shortNameRaw)

	currentRaw := document.Find(".chart__info__sum").Text()
	currentPrice, err2 := getPriceValue(currentRaw)
	if err2 != nil {
		err2 = fmt.Errorf("error parsing the price (%s) for %s: %w", currentRaw, ticker.Name.Full, err2)
		log.Debug(err2)
		chErr <- err2
		return
	}
	ticker.CurrentPrice = currentPrice

	var forecasts []forecast
	forecastsCount := document.Find(".js-review .item__review__sum").
		Each(func(_ int, selection *goquery.Selection) {
			forecastRaw := selection.Text()
			priceValue, err := getPriceValue(forecastRaw)
			if err != nil {
				err = fmt.Errorf("error parsing a forecast target price (%s) for %s: %w", forecastRaw, ticker.Name.Full, err)
				log.Debug(err)
				chErr <- err
				return
			}

			percent := (priceValue - currentPrice) / currentPrice * 100
			forecast := forecast{ExpectedDiff: percent}
			forecasts = append(forecasts, forecast)
		}).Length()

	if forecastsCount > expectedForecastsCount {
		log.Warnf("expected %d forecasts, but got %d for %s", expectedForecastsCount, forecastsCount, ticker.Name.Full)
	}

	datesCount := document.Find(".js-review .item__review__date_big").
		Each(func(i int, selection *goquery.Selection) {
			if i > forecastsCount {
				err := fmt.Errorf("too many time values for %d forecasts for %s", forecastsCount, ticker.Name.Full)
				log.Debug(err)
				chErr <- err
				return
			}

			timeRaw := selection.Text()
			time, err := parseTime(timeRaw)
			if err != nil {
				err = fmt.Errorf("error parsing the time (%s) for %s: %w", timeRaw, ticker.Name.Full, err)
				log.Debug(err)
				chErr <- err
				return
			}

			forecasts[i].Time = time
		}).Length()
	ticker.Forecasts = &forecasts

	if forecastsCount != datesCount {
		err := fmt.Errorf("dates count %d is differ from forecasts count %d for %s", datesCount, forecastsCount, ticker.Name.Full)
		log.Debug(err)
		chErr <- err
		return
	}

	log.Debugf("got forecasts for %s: %v", ticker.Name.Full, ticker.Forecasts)

	chData <- ticker
	log.Debugf("parsing finished. %d forecasts processed for %s", forecastsCount, ticker.Name.Full)
}

func getResponse(url string) (*http.Response, error) {
	log.Debugf("loading content of %s ...", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	log.Debugf("got response from %s", url)

	if resp.StatusCode != 200 {
		err = fmt.Errorf("cannot get document (%s): status code is %d", url, resp.StatusCode)
		log.Debug(err)
		return nil, err
	}

	return resp, nil
}

func closeReader(response *http.Response) error {
	return response.Body.Close()
}

func getUtf8Reader(response *http.Response) (io.Reader, error) {
	contentType := response.Header.Get("Content-Type")
	reader, err := charset.NewReader(response.Body, contentType)
	if err != nil {
		return nil, fmt.Errorf("cannot convert document to utf-8: %w", err)
	}

	return reader, nil
}

// parseOnline runs pages parsing in goroutines, compiles, sorts and returns satellites array.
func parseOnline() (*[]stockTicker, []error) {
	ch, chErr, chQuit := make(chan stockTicker), make(chan error), make(chan int)
	ongoing := 0

	//for _, url := range getProperties().Parser.Urls {
	go parseOnlinePage(getProperties().Parser.URL, ch, chErr, chQuit)
	//}

	var tickers []stockTicker
	var errorz []error

WaiterLoop:
	for {
		select {
		case receivedSat := <-ch:
			tickers = append(tickers, receivedSat)
		case receivedErr := <-chErr:
			errorz = append(errorz, receivedErr)
		case count := <-chQuit:
			ongoing += count
			if ongoing == 0 {
				break WaiterLoop
			}
		}
	}
	close(ch)
	close(chErr)
	close(chQuit)

	return &tickers, errorz
}

func parseOnlinePage(url string, chData chan stockTicker, chErr chan error, chCounter chan int) {
	chCounter <- 1
	defer func() {
		chCounter <- -1
	}()

	httpResponse, err1 := getResponse(url)
	if err1 != nil {
		chErr <- err1
		return
	}
	defer func() {
		if err := closeReader(httpResponse); err != nil {
			chErr <- err
		}
	}()

	reader, err2 := getUtf8Reader(httpResponse)
	if err2 != nil {
		chErr <- err2
		return
	}

	Parse(url, reader, chData, chErr)
}
