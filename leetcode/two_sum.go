package leetcode

func twoSum(nums []int, target int) []int {
	return twoSumBruteForce(nums, target)
}

func twoSumBruteForce(nums []int, target int) []int {
	for x := 0; x < len(nums)-1; x++ {
		for y := x + 1; y < len(nums); y++ {
			attempt := nums[x] + nums[y]
			if attempt == target {
				return []int{x, y}
			}
		}
	}

	return []int{0, 0}
}

func twoSumCache(nums []int, target int) []int {
	cache := map[int]int{}
	for x := 0; x < len(nums); x++ {
		number := nums[x]
		search := target - number
		if cache_index, ok := cache[search]; ok {
			return []int{cache_index, x}
		} else {
			cache[number] = x
		}
	}
	return []int{0, 0}
}
