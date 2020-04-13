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
func TestObtainListOfIndexes(t *testing.T) {
	// given
	expected := "body"
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(expected))
	}))

	defer func() { testServer.Close() }()

	// when
	result, err := datasources.GetListOfIndexes(testServer.URL)

	// then
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Wrong result of a method. \nExpected: %v, \nActual: %v", expected, result)
	}
}

// should obtain stock data for a given company
func TestObtainStockDataForIndex(t *testing.T) {
	// given
	index := "CCC"
	expected := "body"
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(expected))
	}))
	defer func() { testServer.Close() }()

	// when
	result, err := datasources.GetStockDataForIndex(testServer.URL, index)

	// then
	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Wrong result of a method. \nExpected: %v, \nActual: %v", expected, string(result))
	}
}
