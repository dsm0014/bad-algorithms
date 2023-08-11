package bad

// NFactorial has O(n!) time complexity
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

// NToTheN has O(n^n) time complexity
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

// NtoTheNFactorial has O((n^n)!) time complexity
func NtoTheNFactorial(nums []int, n int) int {
	nToTheN := make([]int, NToTheN(nums, len(nums)))
	return NFactorial(nToTheN)
}
