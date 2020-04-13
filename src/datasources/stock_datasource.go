package datasources

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO dodaj obsługę błędów

type StockDatasource interface {
	GetListOfIndexes(stockUrl string) (string, error)
	GetStockDataForIndex(stockUrl, index string) (string, error)
}

type StockDatasourceImpl struct{}

func (sd *StockDatasourceImpl) GetListOfIndexes(stockUrl string) (string, error) {

	if stockUrl == "" {
		stockUrl = os.Getenv("StockURL")
		stockUrl += "company/stock/list"
	}

	response := makeHttpRequest(stockUrl)

	return response, nil
}

func (sd *StockDatasourceImpl) GetStockDataForIndex(stockUrl, index string) (string, error) {

	if stockUrl == "" {
		stockUrl = os.Getenv("StockURL")
		stockUrl += "/financial-ratios/" + index
	}

	response := makeHttpRequest(stockUrl)

	return response, nil
}

func makeHttpRequest(url string) string {
	var client http.Client
	rawResponse, resErr := client.Get(url)
	if resErr != nil {
		fmt.Errorf("Error: %v", resErr)
	}

	defer rawResponse.Body.Close()

	response, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	return string(response)
}
