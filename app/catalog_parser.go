package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"ticker-parser/app/entities"
)

const errorCatalogFetching = 1001

var (
	catalogGetHandlerPath = "/catalog/fetch"

	catalogBaseUrl  = getProperties().Parser.Catalog.BaseUrl
	catalogPageSize = int(getProperties().Parser.Catalog.PageSize)
)

func catalogGetHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("new request from %s: %s", r.RemoteAddr, r.URL.Path)

	catalogHTTPData, httpError := catalogFetch()
	if catalogHTTPData != nil {
		log.Infof("%d items fetched", catalogHTTPData.ItemsCount)
	}

	response := entities.NewHTTPResponse(catalogHTTPData, httpError, 1, r.URL.Path)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Errorf("cannot encode CatalogHTTPData: %s", err)
	}
}

func catalogFetch() (*entities.CatalogHTTPData, *entities.HTTPError) {
	var items []entities.CatalogItem

	for i := 0; ; i++ {
		newItems, errorDetails := catalogFetchPage(i)
		if errorDetails != nil {
			errors := []entities.HTTPErrorDetails{*errorDetails}
			return nil, entities.WrapErrors("cannot fetch catalog", errorCatalogFetching, errors...)
		}

		items = append(items, newItems...)

		if len(newItems) != catalogPageSize {
			break
		}
	}

	return entities.NewCatalogHTTPData(&items), nil
}

func catalogFetchPage(page int) ([]entities.CatalogItem, *entities.HTTPErrorDetails) {
	log.Debugf("getting page %d of catalog ...", page)

	url, err1 := getCatalogPageUrl(page)
	if err1 != nil {
		return nil, &entities.HTTPErrorDetails{
			Reason:       err1.Error(),
			Message:      "cannot get catalog page url",
			Location:     strconv.Itoa(page),
			LocationType: "page",
			ExtendedHelp: "check your configuration parameter: parser.catalog.baseUrl\n" +
				"current value: " + catalogBaseUrl,
		}
	}

	response, err2 := getResponse(url)
	if err2 != nil {
		return nil, &entities.HTTPErrorDetails{
			Reason:       err2.Error(),
			Message:      "cannot fetch catalog page",
			Location:     url,
			LocationType: "url",
			ExtendedHelp: "check your configuration parameter: parser.catalog.baseUrl\n" +
				"current value: " + catalogBaseUrl,
		}
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Error(fmt.Errorf("cannot close response from (%s) body: %w", url, err))
		}
	}()

	var items []entities.CatalogItem
	if err3 := json.NewDecoder(response.Body).Decode(&items); err3 != nil {
		return nil, &entities.HTTPErrorDetails{
			Reason:       err3.Error(),
			Message:      "cannot parse catalog page",
			Location:     url,
			LocationType: "url",
			ExtendedHelp: "check your configuration parameter: parser.catalog.baseUrl\n" +
				"current value: " + catalogBaseUrl,
		}
	}

	log.Debugf("got %d items from %s", len(items), url)
	return items, nil
}

func getCatalogPageUrl(page int) (string, error) {
	req, err := http.NewRequest("GET", catalogBaseUrl, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("sort", "leaders") // blue_chips, leaders, forecast (best forecasts)
	q.Add("type", "share")   // share (stocks), bond, currency
	q.Add("offset", strconv.Itoa(catalogPageSize*page))
	q.Add("limit", strconv.Itoa(catalogPageSize))

	req.URL.RawQuery = q.Encode()

	return req.URL.String(), nil
}
