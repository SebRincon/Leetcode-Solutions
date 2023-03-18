# 53. Maximum Subarray
Given an integer array nums, find the
subarray
with the largest sum, and return its sum.

 

Example 1:

    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:

    Input: nums = [1]
    Output: 1
    Explanation: The subarray [1] has the largest sum 1.

Example 3:

    Input: nums = [5,4,-1,7,8]
    Output: 23
    Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

 

Constraints:

- 1 <= nums.length <= 10<sup>5</sup>
- -104 <= nums[i] <= 10<sup>4</sup>

 

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

---

# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
My first thought was to create a loop and try to calculate (n-1 + n +n-1), then things got too wierd to continue. I then looked up a better approach.
# Approach
<!-- Describe your approach to solving the problem. -->
The approach I found was to loop through the array and continuously add the num[i] to the current sum, then compare that value with the maxSum and set the max to maxSum.

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
O(n)

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
O(n)

# Code
```
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
```