func containsDuplicate(nums []int) bool {
    indexMap := make(map[int]int)
    hasDuplicate := false
    for currentIndex, currNum := range nums{
        if _, isDuplicate := indexMap[currNum]; isDuplicate {
            hasDuplicate = true
        }
        indexMap[currNum] = currentIndex
    }
    return hasDuplicate
}