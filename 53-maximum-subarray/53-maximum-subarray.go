func maxSubArray(nums []int) int {
    maxSub := nums[0]
    curSum := 0
    
    for _, num := range nums {
        if curSum < 0{
            curSum = 0
        }
        curSum += num
        if curSum > maxSub{
            maxSub = curSum
        }
        
    }
    return maxSub
}