package datasources

import (
	"io/ioutil"
	"net/http"
	"os"
)

// StockDatasource interface
type StockDatasource interface {
	GetListOfIndexes(stockUrl string) (string, error)
	GetStockDataForIndex(stockUrl, index string) (string, error)
}

// StockDatasourceImpl implementation of StockDatasource
type StockDatasourceImpl struct{}

// GetListOfIndexes returns JSON with list of indexes
func (sd *StockDatasourceImpl) GetListOfIndexes(stockUrl string) (string, error) {

	if stockUrl == "" {
		stockUrl = os.Getenv("StockURL")
		stockUrl += "company/stock/list"
	}

	response, err := makeHttpRequest(stockUrl)
	if err != nil {
		return "", err
	}

	return response, nil
}

// GetStockDataForIndex ...
func (sd *StockDatasourceImpl) GetStockDataForIndex(stockUrl, index string) (string, error) {

	if stockUrl == "" {
		stockUrl = os.Getenv("StockURL")
		stockUrl += "/financial-ratios/" + index
	}

	response, err := makeHttpRequest(stockUrl)
	if err != nil {
		return "", err
	}

	return response, nil
}

func makeHttpRequest(url string) (string, error) {
	var client http.Client
	rawResponse, resErr := client.Get(url)
	if resErr != nil {
		return "", resErr
	}

	defer rawResponse.Body.Close()

	response, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return "", err
	}

	return string(response), nil
}
