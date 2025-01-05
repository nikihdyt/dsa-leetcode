func searchInsert(nums []int, target int) int {
    low := 0
	high := len(nums) - 1

	for low <= high{
		middle := low + ((high-low)/2)

		if nums[middle] == target {
			return middle
		}

		if target < nums[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	
	return low
}