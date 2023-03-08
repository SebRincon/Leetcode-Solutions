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