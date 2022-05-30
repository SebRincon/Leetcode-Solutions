func maxProfit(prices []int) int {
	// Saftey

    if len(prices) <= 1 {
		return 0
	} 
    
    min, maxSale := prices[0], 0
	// Loop and compare the lowest price, as well as compare the next value to sell

	for _,price := range prices {
		if price < min {
			min = price
        } else if (price-min) > maxSale{
			maxSale = price-min
		}
	}
	
	return  maxSale
}

