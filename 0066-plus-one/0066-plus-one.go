func plusOne(digits []int) []int {
    if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}
	
	for i := 0; i < len(digits); i++ { // digits[last] == 9  // i = index
		observedIdx := len(digits) - i - 1 // looping dari belakang ke depan

		digits[observedIdx]++ // digits[last_idx] pasti akan 10

		if digits[observedIdx] == 10 && observedIdx != 0 {
			digits[observedIdx] = 0
		} else {
			break
		}
	}

	if digits[0] == 10 {
		digits = append([]int{1}, digits...)
		digits[1] = 0
	}

	return digits
}