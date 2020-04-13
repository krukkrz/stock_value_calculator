package test

import (
	"io/ioutil"
	"os"
	"reflect"
	"stockvaluecalculator/src/models/entities"
	"stockvaluecalculator/src/repositories"
	"stockvaluecalculator/test/mocks"
	"testing"
)

// should build an Indexes array for a given json
func TestGetIndexesArray(t *testing.T) {
	// given
	response := buildMockJSONResponseFromFile("./mocks/json_responses/index_list.json", t)
	mockDatasource := mocks.NewMockDatasource(string(response))
	repository := repositories.NewStockRepositoryImpl(mockDatasource)
	expected := []string{"SPY", "CMCSA", "KMI", "INTC", "MU", "GDX"}

	// when
	indexes := repository.GetIndexesArray()

	// then
	compare := reflect.DeepEqual(indexes, expected)
	if !compare {
		t.Errorf("Wrong method result. \nExpected: %v \nActual %v", expected, indexes)
	}
}

// should build an Company model from a given json
func TestGetCompanyForIndex(t *testing.T) {
	// given
	response := buildMockJSONResponseFromFile("./mocks/json_responses/company_profile.json", t)
	mockDatasource := mocks.NewMockDatasource(string(response))
	repository := repositories.NewStockRepositoryImpl(mockDatasource)
	expected := &entities.Company{
		PriceEarningsRatio:         -23.28911282026003,
		PriceEarningsToGrowthRatio: -5137.711820504985,
		DividendYield:              0,
	}

	// when
	company := repository.GetCompanyForIndex("CCC")

	// then
	if !reflect.DeepEqual(company, expected) {
		t.Errorf("Wrong method result. \nExpected: %v \nActual %v", expected, company)
	}
}

func buildMockJSONResponseFromFile(path string, t *testing.T) []byte {
	expectedJSON, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	read, err := ioutil.ReadAll(expectedJSON)
	if err != nil {
		t.Error(err)
	}

	return read
}
