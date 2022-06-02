func productExceptSelf(nums []int) []int {
	// Create array of ones
	arrayLen := len(nums)
	_result := make([]int, arrayLen)
	for i := range _result {
		_result[i] = 1
	}

	// Work forwards incrementing the product
	prefix := 1
	for i, num := range nums {
		_result[i] = prefix
		prefix *= num
	}

	// Working backwards and incremeting the postfix product
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		_result[i] *= postfix
		postfix *= nums[i]
	}
	return _result

}
