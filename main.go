package main

import (
	"os"
	"stockvaluecalculator/src/utils"
)

func main() {
	utils.SetEnvVariables()

	println(os.Getenv("StockURL"))
}
