package main

// Time Complexity: O(n^2)
// Space Complexity: O(1)
import "fmt"

func main() {
	test := []int{1, 2, 4, 6}
	fmt.Print(twoSum(test, 10))
	fmt.Print(twoSum_faster(test, 10))
}

func twoSum(nums []int, target int) []int {
	arrayLen := len(nums)
	_answer := []int{}

	for i := 0; i < arrayLen; i++ {
		_index := nums[i]
		for f := i + 1; f < arrayLen; f++ {
			if _index+nums[f] == target {
				return append(_answer, []int{i, f}...)
			}
		}
	}
	return _answer

}
