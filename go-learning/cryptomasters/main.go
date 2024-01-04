package main

import (
	"fmt"
	"sync"

	"github.com/sultan/crypto/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}

	wg.Wait()

	// go getCurrencyData("BTC")
	// go getCurrencyData("ETH")
	// go getCurrencyData("BCH")
	//
	// time.Sleep(time.Second * 5)

	// fmt.Println(rate.Price, err)
	// fmt.Println()
	// print(rate.Price, err)
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)

	}
}
