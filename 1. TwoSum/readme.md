# 1. TwoSum Description
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

___

Example 1:

    Input: nums = [2,7,11,15], target = 9
    Output: [0,1]
    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

    Input: nums = [3,2,4], target = 6
    Output: [1,2]

Example 3:

    Input: nums = [3,3], target = 6
    Output: [0,1]

 

Constraints:

    2 <= nums.length <= 104
    -109 <= nums[i] <= 109
    -109 <= target <= 109
    Only one valid answer exists.



# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
I started with the brute force method, iterating over the list n*n times, I later came back using the hashmap method I learned in another solution.


# Approach
<!-- Describe your approach to solving the problem. -->
Loop through the array and store the index & values in a hashmap, then attempt to locate the current index - target number in the hashmap and return the index if succesful or store the value if not.


# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
0(n)

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
0(n)

# Code
```

func twoSum(nums []int, target int) []int {
    // Hashmap that hold all values
    s := make(map[int]int)

    // Loop through values
    for idx, num := range nums {
        // if the target - the current index number exists in the map return the pair
        if pos, ok := s[target-num]; ok {
            return []int{pos, idx}
        }
        s[num] = idx
    }
    return []int{}
}



//  ┌─────────────┐
//  │Target: 4    │
//  │Nums: 1,3,5,2│
//  └─────────────┘
//
// ┌───────┐
// │Loop 1 │──────────┐
// └───────┘          ▼
//  ┌──────────────────────────────────┐
//  │is [4-1](target, num) in the map ?│
//  └──────────────────────────────────┘
//                    │
//     Yes────────────┴───No
//       │                │
//       ▼                ▼
//  ┌───────────┐ Map─────┬───┬───┬───┐
//  │  Return   │   │Value│ 1 │   │   │
//  │{idx, num} │   ├─────┼───┼───┼───┤
//  └───────────┘   │Index│ 0 │   │   │
//                  └─────┴───┴───┴───┘
//                          │
//     ┌────────────────────┘
//     │
//     ▼
// ┌───────┐
// │Loop n │───────────┐
// └───────┘           ▼
//   ┌──────────────────────────────────┐
//   │is [4-3](target, num) in the map ?│
//   └──────────────────────────────────┘
//                     │
//      Yes────────────┴───No
//        │                │
//        ▼                ▼
//   ┌───────────┐ Map─────┬───┬───┬───┐
//   │  Return   │   │Value│ 1 │ 3 │   │
//   │{idx, num} │   ├─────┼───┼───┼───┤
//   └───────────┘   │Index│ 0 │ 1 │   │
//         │         └─────┴───┴───┴───┘
//         │
//         │  ┌─────────────────┐
//         └─▶│ Returns {0, 1}  │
//            └─────────────────┘
```