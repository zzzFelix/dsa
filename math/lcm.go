package math

// least common multiple
func LCM(a, b int, nums ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}

	return result
}
