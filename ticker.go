package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"time"
)

//nolint GoUnusedType
type apiResponse struct {
	APIVersion string      `json:"apiVersion"`
	Method     string      `json:"method"`
	Data       interface{} `json:"data"`
	Error      error       `json:"error"`
}

//nolint GoUnusedType
type apiError struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

type tickerCollection struct {
	Tickers *[]stockTicker `json:"tickers"`
}

type stockTicker struct {
	Name         tickerName  `json:"name"`
	CurrentPrice float64     `json:"price"`
	Consensus    float64     `json:"consensus"`
	Forecasts    *[]forecast `json:"forecasts"`
}

type tickerName struct {
	Full  string `json:"full"`
	Short string `json:"short"`
}

type forecast struct {
	ExpectedDiff float64   `json:"expectedDiff"`
	Time         time.Time `json:"time"`
}

//type JSONTime time.Time
//
//func (t JSONTime) MarshalJSON() ([]byte, error) {
//	time.Time(t).Format("")
//	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
//	return []byte(stamp), nil
//}

func (ptr *forecast) String() string {
	return fmt.Sprintf("%+f %s", ptr.ExpectedDiff, ptr.Time.String())
}

func filterOldForecasts(ticker *stockTicker) error {
	/* Filter old forecasts */
	var newForecasts []forecast
	var count int
	for _, forecast := range *ticker.Forecasts {
		if forecast.Time.Before(time.Now().AddDate(0, -1, 0)) {
			continue
		}
		newForecasts = append(newForecasts, forecast)
		count++
	}
	ticker.Forecasts = &newForecasts

	if count < 5 {
		return fmt.Errorf("not enough actual forecasts for ticker %s", ticker.Name)
	}
	return nil

}

//noinspection GoNilness
func filterExtremeForecasts(ticker *stockTicker) error {
	/* Filter extreme values */
	threshold := getProperties().Filters.ExtremeValues.Threshold
	logrus.Debugf("threshold to exclude extreme value is %f", threshold)

	var prices []float64
	for _, forecast := range *ticker.Forecasts {
		prices = append(prices, forecast.ExpectedDiff)
	}
	sort.Float64s(prices)

	count := len(prices)
	if count < 2 {
		return nil
	}

	if (prices[1] - prices[0]) > threshold {
		newForecasts, err1 := removeForecast(ticker.Forecasts, prices[0])
		if err1 != nil {
			return err1
		}
		ticker.Forecasts = newForecasts
	}

	if (prices[count-1] - prices[count-2]) > threshold {
		newForecasts, err1 := removeForecast(ticker.Forecasts, prices[count-1])
		if err1 != nil {
			return err1
		}
		ticker.Forecasts = newForecasts
	}

	return nil
}

func removeForecast(forecastsPtr *[]forecast, value float64) (*[]forecast, error) {
	forecasts := *forecastsPtr
	for i, forecast := range forecasts {
		if forecast.ExpectedDiff == value {
			last := len(forecasts) - 1
			forecasts[i] = forecasts[last]
			newForecasts := forecasts[:last]
			return &newForecasts, nil
		}
	}
	return nil, fmt.Errorf("cannot find value %f to remove forecast", value)
}
