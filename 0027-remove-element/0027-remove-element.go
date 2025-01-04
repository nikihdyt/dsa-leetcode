func removeElement(nums []int, val int) int {
	counter := 0

	for _, num := range nums {
		if num == val {
			continue
		}
        nums[counter] = num	
        counter++	
	} 

	return counter  
}