package mocks

type MockDatasource struct {
	expected string
}

func NewMockDatasource(expected string) *MockDatasource {
	return &MockDatasource{
		expected: expected,
	}
}

func (md *MockDatasource) GetListOfIndexes(stockUrl string) (string, error) {
	return md.expected, nil
}
func (md *MockDatasource) GetStockDataForIndex(stockUrl, index string) (string, error) {
	return md.expected, nil
}
