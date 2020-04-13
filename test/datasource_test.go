package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"stockvaluecalculator/src/datasources"
	"stockvaluecalculator/src/utils"
	"testing"
)

func TestMain(t *testing.M) {
	utils.SetEnvVariables()
	// utils.SetTestEnvVariables()

	os.Exit(t.Run())
}

// Should receive http response body with a list of indexes as JSON object when response code is 200
func TestGetListOfIndexes(t *testing.T) {
	// given
	datasource := datasources.StockDatasourceImpl{}
	expected := "body"
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(expected))
	}))

	defer func() { testServer.Close() }()

	// when
	result, err := datasource.GetListOfIndexes(testServer.URL)

	// then
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Wrong result of a method. \nExpected: %v, \nActual: %v", expected, result)
	}
}

// should obtain stock data for a given company
func TestGetStockDataForIndex(t *testing.T) {
	// given
	datasource := datasources.StockDatasourceImpl{}

	index := "CCC"
	expected := "body"
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(expected))
	}))
	defer func() { testServer.Close() }()

	// when
	result, err := datasource.GetStockDataForIndex(testServer.URL, index)

	// then
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Wrong result of a method. \nExpected: %v, \nActual: %v", expected, string(result))
	}
}
