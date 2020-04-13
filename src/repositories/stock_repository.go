package repositories

import (
	"encoding/json"
	"fmt"
	ds "stockvaluecalculator/src/datasources"
	"stockvaluecalculator/src/models/dto"
	"stockvaluecalculator/src/models/dto/companydto"
	"stockvaluecalculator/src/models/entities"
	"strconv"
)

type StockRepository interface {
	GetIndexesArray() []string
}

type StockRepositoryImpl struct {
	datasource ds.StockDatasource
}

func NewStockRepositoryImpl(datasource ds.StockDatasource) *StockRepositoryImpl {
	return &StockRepositoryImpl{
		datasource: datasource,
	}
}

// TODO obsłuż błędy poprawnie

func (sr *StockRepositoryImpl) GetIndexesArray() []string {
	jsonString, err := sr.datasource.GetListOfIndexes("")
	if err != nil {
		fmt.Errorf("ERROR: %v", err)
	}

	var indexes dto.Indexes
	err = json.Unmarshal([]byte(jsonString), &indexes)
	if err != nil {
		fmt.Errorf("ERROR: %v", err)
	}

	var arr []string

	for i := 0; i < len(indexes.SymbolsList); i++ {
		arr = append(arr, indexes.SymbolsList[i].Symbol)
	}

	return arr
}

func (sr *StockRepositoryImpl) GetCompanyForIndex(index string) *entities.Company {

	jsonString, err := sr.datasource.GetStockDataForIndex("", index)
	if err != nil {
		fmt.Errorf("ERROR: %v", err)
	}

	var companyDto companydto.Company
	err = json.Unmarshal([]byte(jsonString), &companyDto)
	if err != nil {
		fmt.Errorf("ERROR: %v", err)
	}

	priceEarningsToGrowthRatio, _ := strconv.ParseFloat(companyDto.Ratios[0].InvestmentValuationRatios.PriceEarningsToGrowthRatio, 64)
	dividendYield, _ := strconv.ParseFloat(companyDto.Ratios[0].InvestmentValuationRatios.DividendYield, 64)

	// TODO przemyśl, czy nie użyć tych wartości z trzech dni uśrednionych
	return &entities.Company{
		PriceEarningsToGrowthRatio: priceEarningsToGrowthRatio,
		DividendYield:              dividendYield,
	}
}
