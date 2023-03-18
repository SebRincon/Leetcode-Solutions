// Helper max function
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    var maxSub = nums[0]
    var curSum = 0
    // Loop through numbers
    for _, num := range nums {
        // Calculate if the current sum range has become negative
        // If negative reset to zero
        if curSum < 0 {
            curSum = 0
        }
        // Add to the current sum range
        curSum += num
        // Calcuate if the current sum is larger than the max sum
        maxSub = max(maxSub, curSum)
    }

    return maxSub

    
}

// ┌─────────────────────────────┐
// │nums: [-2,1,-3,4,-1,2,1,-5,4]│
// └─────────────────────────────┘
// ┌──────────┐
// │maxSum: 0 │
// │curSum: 0 │
// └──────────┘
// ┌───────┐
// │n Loops│─────┐
// └───────┘     ▼
//    ┌────────────────────┐
//    │ curSum += nums[i]  │
//    └────────────────────┘
//               │
//             Is curSum < 0 ?
//               │       ┌─────────────┐
//               ├─Yes──▶│ curSum = 0  │
//               No      └─────────────┘
//               │              │
//               ▼              │
//    ┌────────────────────┐    │
//    │ curSum += nums[i]  │◀───┘
//    └────────────────────┘
//               │
//               └─────┐
//                     ▼
//    ┌────────────────────────────────┐
//    │  maxSum = max(curSum, maxSum)  │
//    └────────────────────────────────┘