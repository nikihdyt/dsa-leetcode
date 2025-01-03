func twoSum(nums []int, target int) []int {
	mMap := make(map[int]int)

	for idx, i := range nums {
		complement := target - i
		val, ok := mMap[complement]
		if ok {
			return []int{idx, val}
		} else {
			mMap[i] = idx
		}
	}

    return []int{0, 0}
}