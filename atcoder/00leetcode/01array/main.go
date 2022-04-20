package array

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for ni, n := range nums {
		for ti, t := range nums {
			if ni == ti {
				continue
			}
			if n+t == target {
				res[0] = ni
				res[1] = ti
				break
			}
		}
	}

	return res
}
