func twoSum(nums []int, target int) []int {
	arrayLen := len(nums) // result slice
	_answer := []int{}
    
    // Loop each number and test the rest of the array on the that number
	for i := 0; i < arrayLen; i++ {
		_index := nums[i]
		for f := i + 1; f < arrayLen; f++ {
			if _index+nums[f] == target {
                // append to the result
				return append(_answer, []int{i, f}...)
			}
		}
	}
	return _answer

}