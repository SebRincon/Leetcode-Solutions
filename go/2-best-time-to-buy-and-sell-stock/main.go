package main

func main() {
	prices := []int{2, 4, 1}

	println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	// Saftey
	if len(prices) <= 1 {
		return 0
	}

	minPrice := 1000
	maxSale := 0
	// Loop and compare the lowest price, as well as compare the next value to sell
	for _, num := range prices {
		if num < minPrice {
			minPrice = num
		} else if (num - minPrice) > maxSale {
			maxSale = minPrice - num

		}

	}

	return maxSale

}
