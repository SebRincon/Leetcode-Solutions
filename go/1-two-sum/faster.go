package main

func twoSum_faster(nums []int, target int) []int {
	// Create a map{intKey:intValue}
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		// test if current index has an available pair in the map
		// if so return the index of the pair
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		// otherwise add current index to map storage
		indexMap[currNum] = currIndex
	}
	return []int{}
}
