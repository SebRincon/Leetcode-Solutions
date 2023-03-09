# 121. BestTrade
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
___

Example 1:

    Input: prices = [7,1,5,3,6,4]
    Output: 5
    Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
    Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:

    Input: prices = [7,6,4,3,1]
    Output: 0
    Explanation: In this case, no transactions are done and the max profit = 0.

    

Constraints:

    1 <= prices.length <= 10 <sup>5<sup>
    0 <= prices[i] <= 10 <sup>4<sup>



# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
I thought of a brute force comparing every value with its successors, but this would be very costly, so I found the solution using two pointers.

# Approach
<!-- Describe your approach to solving the problem. -->
Using two pointers we can iterate over the array and compare values if the meet a profit making senario, otherwise set the new low.

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
O(n)

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
O(n)

# Code
```
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
// └─────────────────────┘   ┌──────────────────────────────┐
//                           │l = left pointer (low price)  │
//                           │r = right pointer (high price)│
// ┌───────┐                 └──────────────────────────────┘
// │n Loops│─────────┐                       ▲
// └───────┘         ▼                ┌──────┘
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

```