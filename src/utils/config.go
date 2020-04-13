package utils

import "os"

func SetEnvVariables() {
	os.Setenv("StockURL", "https://financialmodelingprep.com/api/v3/")
}
