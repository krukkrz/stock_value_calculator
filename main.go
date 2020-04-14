package main

import (
	"fmt"
	"stockvaluecalculator/src/datasources"
	"stockvaluecalculator/src/repositories"
	"stockvaluecalculator/src/usecases"
	"stockvaluecalculator/src/utils"
	"time"
)

func main() {
	start := time.Now()

	utils.SetEnvVariables()
	datasource := datasources.StockDatasourceImpl{}
	repository := repositories.NewStockRepositoryImpl(&datasource)

	println("Running a calculation...")
	result := usecases.CalculateValuesForCurrentIndexes(repository)
	println("========== RESULTS ===========")
	for index, value := range result {
		fmt.Printf("| %v | %v | \n", index, value)
	}

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %v \nTotal number of indexes: %v \n", elapsed, len(result))
}
