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
