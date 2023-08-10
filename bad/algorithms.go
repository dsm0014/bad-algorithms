package bad

func NFactorial(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	depth := 0
	for i := 0; i < len(nums); i++ {
		depth += NFactorial(nums[1:])
	}
	return depth
}

func NToTheN(nums []int, n int) int {
	if n == 0 {
		return 1
	}
	depth := 0
	for i := 0; i < len(nums); i++ {
		depth += NToTheN(nums, n-1)
	}
	return depth
}
