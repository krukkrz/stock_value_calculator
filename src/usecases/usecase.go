package usecases

import (
	"fmt"
	"stockvaluecalculator/src/models/entities"
	"stockvaluecalculator/src/repositories"
)

// CalculateValuesForCurrentIndexes calculates values for available indexes
func CalculateValuesForCurrentIndexes(repository repositories.StockRepository) map[string]float64 {
	result := map[string]float64{}

	indexes := repository.GetIndexesArray()
	println("Obtaining company data for each index...")

	// TODO implement concurrency here
	for _, index := range indexes {
		company := repository.GetCompanyForIndex(index)
		if !isCompanyEmpty(company) {
			value := calculateValue(company)
			result[index] = value
			fmt.Printf("Done for index: %v \n", index)
		}
	}

	return result
}

func calculateValue(company *entities.Company) float64 {
	return (company.PriceEarningsRatio) / ((company.PriceEarningsToGrowthRatio / company.PriceEarningsRatio) + company.DividendYield)
}

func isCompanyEmpty(company *entities.Company) bool {
	dy := company.DividendYield == 0
	pe := company.PriceEarningsRatio == 0
	peg := company.PriceEarningsToGrowthRatio == 0

	if dy && pe && peg {
		return true
	}
	return false
}
