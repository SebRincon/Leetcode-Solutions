# 217. Contains Duplicate


Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 
Example 1:

    Input: nums = [1,2,3,1]
    Output: true

Example 2:

    Input: nums = [1,2,3,4]
    Output: false

Example 3:

    Input: nums = [1,1,1,3,3,4,3,2,4,2]
    Output: true

 

Constraints:

- 1 <= nums.length <= 10<sup>5</sup>
- -109 <= nums[i] <= 10<sup>9</sup>
---



# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
My first thought was to use the map to compare the current value against.

# Approach
<!-- Describe your approach to solving the problem. -->
I found out that go has a validation feature when indexing the map, so I used that to compare against the current value.

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->
O(n)

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->
O(n)

# Code
```
func containsDuplicate(nums []int) bool {
    values := make(map[int]int, len(nums))

    // Loop over values
    for idx, num := range nums {
        // Check if current value in the map
        _, found := values[num]

        if found{
            return true
        }
        // Insert the value into the map if a dup is not found
        values[num] = idx
    }

    // return false at the end
    return false
}

//    ┌───────┐
// ┌─▶│n Loops│──────────┐
// │  └───────┘          ▼
// │   ┌──────────────────────────────────┐
// │   │   is current value is in map ?   │
// │   └──────────────────────────────────┘
// │                     │
// │                     │
// │          No─────────┴─────Yes
// │          │                 │
// │          ▼                 ▼
// │   ┌─────────────┐   ┌─────────────┐
// │   │ insert the  │   │ return true │
// │   │ current val │   └─────────────┘
// │   │  into map   │
// │   └─────────────┘
// │          │
// │          │
// └──────────┘
```