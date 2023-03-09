func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func maxProfit(prices []int) int {
    // Left pointer as low price target 
    var l = 0
    var maxP = 0

    // Right pointer as high price target 
    for r, _ := range prices {
        var profit = 0
        // Check if the left index is lower than the right
        if (prices[l] < prices[r]){
            // Compare profits
            profit = prices[r] - prices[l]
            maxP = max(profit, maxP)
        }else {
            // Set new low target
            l = r
        }
        
        r++
    }
    // Continue right pointer
    return maxP
}


// ┌─────────────────────┐
// │prices: [7,1,5,3,6,4]│
// └─────────────────────┘   ┌─────────────────────────────┐
//                           │l = left pointer (low price) │
//                           │r = right pointer (high      │
// ┌───────┐                 └─────────────────────────────┘
// │n Loops│─────────┐                      ▲
// └───────┘         ▼                ┌─────┘
// ┌──────────────────────────────────?
// │    is prices[l] < prices[r] ?    │
// └──────────────────────────────────┘
//                   │
//        ┌No────────┴───────────Yes
//        │                        │
//        ▼                        ▼
// ┌──────────────┐           ┌─────────┬──────────────────────┐
// │new Low point │           │ Profit  │prices[r] - prices[l] │
// └──────────────┘           ├─────────┼──────────────────────┤
//         │                  │maxProfit│max(profit, maxProfit)│
//         │                  └─────────┴───┬──────────────────┘
//         └─────┬──────────────────────────┘
//               │
//               ▼
//     ┌───────────────────┐
//     │ Return maxProfit  │
//     └───────────────────┘
