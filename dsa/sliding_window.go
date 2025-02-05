package dsa

func MaxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	deque := []int{} // store indices

	for i := 0; i < len(nums); i++ {
		// Remove indices that are out of the current window
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove indices of smaller elements as they are not useful
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add current index
		deque = append(deque, i)

		// Append the maximum to the result list once the first window is complete
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

func LongestSubstringKDistinct(s string, k int) int {
	charCount := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		charCount[s[right]]++

		// Shrink window if distinct character count exceeds k
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		// Update max length of valid substring
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func TumblingWindowSum(nums []int, k int) []int {
	result := []int{}
	n := len(nums)

	for i := 0; i < n; i += k {
		sum := 0
		// Sum all elements in the current window
		for j := i; j < i+k && j < n; j++ {
			sum += nums[j]
		}
		result = append(result, sum)
	}
	return result
}
