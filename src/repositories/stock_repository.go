package repositories

import (
	"encoding/json"
	"fmt"
	"regexp"
	ds "stockvaluecalculator/src/datasources"
	"stockvaluecalculator/src/models/dto"
	"stockvaluecalculator/src/models/dto/companydto"
	"stockvaluecalculator/src/models/entities"
	"strconv"
)

// StockRepository a repository interface
type StockRepository interface {
	GetIndexesArray() []string
	GetCompanyForIndex(index string) *entities.Company
}

// StockRepositoryImpl implementation of an interface
type StockRepositoryImpl struct {
	datasource ds.StockDatasource
}

// NewStockRepositoryImpl constructor for implementation of a stock repository
func NewStockRepositoryImpl(datasource ds.StockDatasource) *StockRepositoryImpl {
	return &StockRepositoryImpl{
		datasource: datasource,
	}
}

// GetIndexesArray returns a slice of strings with available indexes
func (sr *StockRepositoryImpl) GetIndexesArray() []string {
	jsonString, err := sr.datasource.GetListOfIndexes("")
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	var indexes dto.Indexes
	err = json.Unmarshal([]byte(jsonString), &indexes)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	var arr []string

	for i := 0; i < len(indexes.SymbolsList); i++ {
		arr = append(arr, indexes.SymbolsList[i].Symbol)
	}

	return arr
}

// GetCompanyForIndex returns a Company struct with available data for given index
func (sr *StockRepositoryImpl) GetCompanyForIndex(index string) *entities.Company {

	jsonString, err := sr.datasource.GetStockDataForIndex("", index)

	// remove unnecesary comments
	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	cleanJSONBytes := re.ReplaceAll([]byte(jsonString), nil)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	if string(cleanJSONBytes) != "{ }" {
		var companyDto companydto.Company

		err = json.Unmarshal(cleanJSONBytes, &companyDto)

		if err != nil {
			fmt.Printf("ERROR: %v", err)
		}

		priceEarningsRatio, _ := strconv.ParseFloat(companyDto.Ratios[0].InvestmentValuationRatios.PriceEarningsRatio, 64)
		priceEarningsToGrowthRatio, _ := strconv.ParseFloat(companyDto.Ratios[0].InvestmentValuationRatios.PriceEarningsToGrowthRatio, 64)
		dividendYield, _ := strconv.ParseFloat(companyDto.Ratios[0].InvestmentValuationRatios.DividendYield, 64)

		return &entities.Company{
			PriceEarningsRatio:         priceEarningsRatio,
			PriceEarningsToGrowthRatio: priceEarningsToGrowthRatio,
			DividendYield:              dividendYield,
		}
	}
	return &entities.Company{}
}
