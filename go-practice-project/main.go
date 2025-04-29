package main

import (
	"fmt"

	"example.com/m/cmdmanager"
	"example.com/m/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// priceJob := prices.NexTaxIncludedPriceJob(fm, taxRate)
		cmdm := cmdmanager.New()
		priceJob := prices.NexTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
