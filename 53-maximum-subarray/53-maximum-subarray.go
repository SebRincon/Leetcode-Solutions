func maxSubArray(nums []int) int {
    maxSub := nums[0]
    curSum := 0
    
    for _, num := range nums {
        // If sum falls to negative reset count
        if curSum < 0{
            curSum = 0
        }
        // Increment and set maxSum if the current count is larger
        curSum += num
        if curSum > maxSub{
            maxSub = curSum
        }
        
    }
    return maxSub
}