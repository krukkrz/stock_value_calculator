package usecases

import "stockvaluecalculator/src/repositories"

func CalculateValuesForCurrentIndexes(repository repositories.StockRepository) map[string]float64 {
	result := map[string]float64{}
	// bierze listę indeksów
	indexes := repository.GetIndexesArray()

	for _, index := range indexes {
		// dla każdego zwraca obiekt Company
		company := repository.GetCompanyForIndex(index)
		// z company wylicza value
		value := (company.PriceEarningsRatio) / ((company.PriceEarningsToGrowthRatio / company.PriceEarningsRatio) + company.DividendYield)

		// index i value dodaje do mapy
		result[index] = value
	}

	// na koniec zwraca mapę
	return result
}
