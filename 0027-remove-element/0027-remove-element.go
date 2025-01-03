func removeElement(nums []int, val int) int {
	counter := 0

	for _, num := range nums {
		if num != val {
			nums[counter] = num	
			counter++
		}
		fmt.Println(nums)
	}

	return counter  
}