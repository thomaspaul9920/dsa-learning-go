package problems

func GetConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := range n {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}
