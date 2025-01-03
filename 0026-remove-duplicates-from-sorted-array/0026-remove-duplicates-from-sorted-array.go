func removeDuplicates(nums []int) int {
    mMap := make(map[int]bool)
	counter := 0

	for _, num := range nums {
		_, exists := mMap[num]
		if exists { // if exists, next
			continue
		} else { // if not, add ke map
			mMap[num] = true
			nums[counter] = num
			counter++
		}
	}

	return counter
}