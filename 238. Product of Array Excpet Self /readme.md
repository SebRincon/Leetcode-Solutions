# 238. Product of Array Except Self

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

 
---
Example 1:

    Input: nums = [1,2,3,4]
    Output: [24,12,8,6]

Example 2:

    Input: nums = [-1,1,0,-3,3]
    Output: [0,0,9,0,0]

 

Constraints:

2 <= nums.length <= 10<sup>5</sup>
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

 

Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
I thought of looping through the values 3 times, the first two times to get the prefix and postfix values and the third to combine values.
I later found the approach to use a single array and only two loops.

# Approach
<!-- Describe your approach to solving the problem. -->
I used a single array and two loops, the first loop was to find the postfix values for each index, the second loop I went backwards in the array and combined the postfix values with the prefix values.

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
O(n)

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
O(1)

# Code
```
func productExceptSelf(nums []int) []int {
    res := make ([]int, len(nums))

    for i := range res {
        res[i] = 1 
    }


    prefix := 1
    for idx, num := range(nums){
        res[idx] = prefix 
        prefix *= num
    }

    

    postfix := 1
    for i := len(nums)-1; i>=0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
    }

// ┌───────────────────────────────────────────────────────────────┐
// │┌───────────────┐                                              │
// ││nums: [1,2,3,4]│                                              │
// │└┬──────────────┘                                              │
// │ │                                                             │
// │ │ ┌───────────────┐ ┌───────────┐                             │
// │ │ │res: [1,1,1,1] │ │prefix:  1 │                             │
// │ │ │[1] * len(nums)│ └───────────┘                             │
// │ │ └───────────────┘                                           │
// │ │                                                             │
// │ │                                                             │
// │ For each num in nums                                          │
// │ │                                                             │
// │ │ ┌────────────────────────┐                                  │
// │ │ │    res[i] = prefix     │   ┌───────────────┐              │
// │ └▶│      prefix *= 1       │──▶│res: [1,1,1,1] │              │
// │   ├───────────────┬────────┘   └───────────────┘              │
// │   │prefix: 1      │                                           │
// │   └┬──────────────┘                                           │
// │    │ ┌────────────────────────┐                               │
// │    │ │    res[i] = prefix     │   ┌───────────────┐           │
// │    └▶│      prefix *= 2       │──▶│res: [1,1,1,1] │           │
// │      ├───────────────┬────────┘   └───────────────┘           │
// │      │prefix: 2      │                                        │
// │      └┬──────────────┘                                        │
// │       │ ┌────────────────────────┐                            │
// │       │ │    res[i] = prefix     │   ┌───────────────┐        │
// │       └▶│      prefix *= 3       │──▶│res: [1,1,2,1] │        │
// │         ├───────────────┬────────┘   └───────────────┘        │
// │         │prefix: 6      │                                     │
// │         └─┬─────────────┘                                     │
// │           │ ┌────────────────────────┐                        │
// │           │ │    res[i] = prefix     │   ┌───────────────┐    │
// │           └▶│      prefix *= 4       │──▶│res: [1,1,2,6] │    │
// │             ├───────────────┬────────┘   └───────────────┘    │
// │             │prefix: 24     │                                 │
// │             └───────────────┘                                 │
// └──────┬────────────────────────────────────────────────────────┘
//        │
//        │   ┌───────────────┐
//        └──▶│res: [1,1,2,6] │
//            └───────────────┘
//
// ┌───────────────────────────────────────────────────────────────┐
// │ ┌───────────────┐ ┌───────────┐                               │
// │ │nums: [4,3,2,1]│ │postfix:  1│                               │
// │ └┬──────────────┘ └───────────┘                               │
// │  │                                                            │
// │  For each num in nums                                         │
// │  Looping in a reverse                                         │
// │  │                                                            │
// │  │ ┌────────────────────────┐                                 │
// │  │ │   res[-1] *= posfix    │    ┌───────────────┐            │
// │  └▶│      posfix *= 4       │──▶ │res: [1,1,2,6] │            │
// │    ├───────────────┬────────┘    └───────────────┘            │
// │    │postfix: 4     │                                          │
// │    └┬──────────────┘                                          │
// │     │ ┌────────────────────────┐                              │
// │     │ │   res[-2] *= posfix    │   ┌───────────────┐          │
// │     └▶│      posfix *= 3       │──▶│res: [1,1,8,6] │          │
// │       ├───────────────┬────────┘   └───────────────┘          │
// │       │postfix: 12    │                                       │
// │       └┬──────────────┘                                       │
// │        │ ┌────────────────────────┐                           │
// │        │ │   res[-3] *= posfix    │    ┌───────────────┐      │
// │        └▶│      posfix *= 2       │───▶│res: [1,12,8,6]│      │
// │          ├───────────────┬────────┘    └───────────────┘      │
// │          │postfix: 24    │                                    │
// │          └─┬─────────────┘                                    │
// │            │ ┌────────────────────────┐                       │
// │            │ │    res[i] = posfix     │   ┌────────────────┐  │
// │            └▶│      posfix *= 1       │──▶│res: [24,12,8,6]│  │
// │              ├───────────────┬────────┘   └────────────────┘  │
// │              │postfix: 24    │                                │
// │              └───────────────┘                                │
// │                                                               │
// └────────┬──────────────────────────────────────────────────────┘
//          │
//          │   ┌────────────────┐
//          └──▶│res: [24,12,8,6]│
//              └────────────────┘
```