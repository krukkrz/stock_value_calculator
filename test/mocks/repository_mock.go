package mocks

import "stockvaluecalculator/src/models/entities"

type MockStockRepository struct {
}

func NewMockStockRepository() *MockStockRepository {
	return &MockStockRepository{}
}

func (msr *MockStockRepository) GetIndexesArray() []string {
	return []string{"SPY", "CMCSA", "KMI", "INTC", "MU", "GDX"}
}

func (msr *MockStockRepository) GetCompanyForIndex(index string) *entities.Company {
	return &entities.Company{
		PriceEarningsRatio:         2,
		PriceEarningsToGrowthRatio: 2,
		DividendYield:              0,
	}
}
