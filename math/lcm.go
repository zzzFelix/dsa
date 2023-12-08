package math

import "errors"

// least common multiple
func LCM(nums ...int) (int, error) {
	if len(nums) < 2 {
		return -1, errors.New("at least two numbers are required to find LCM")
	}

	result := nums[0] * nums[1] / GCD(nums[0], nums[1])

	for i := 0; i < len(nums[2:]); i++ {
		resultTmp, err := LCM(result, nums[i])
		result = resultTmp
		if err != nil {
			return -1, err
		}
	}

	return result, nil
}
