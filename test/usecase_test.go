package test

import (
	"reflect"
	"stockvaluecalculator/src/usecases"
	"stockvaluecalculator/test/mocks"
	"testing"
)

// should return an object  (map? struct) with a list of index + value pairs
func TestCalculateValuesForCurrentIndexes(t *testing.T) {
	// given
	mockRepository := mocks.NewMockStockRepository()
	expected := map[string]float64{
		"SPY":   2,
		"CMCSA": 2,
		"KMI":   2,
		"INTC":  2,
		"MU":    2,
		"GDX":   2,
	}

	// when
	usecase := usecases.CalculateValuesForCurrentIndexes(mockRepository)

	// then
	if !reflect.DeepEqual(usecase, expected) {
		t.Errorf("Wrong result of a method. \nExpected: %v, \nActual: %v", expected, usecase)
	}
}
