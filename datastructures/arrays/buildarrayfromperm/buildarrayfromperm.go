package buildarrayfromperm

func buildArray(nums []int) (ans []int) {
	for idx := range nums {
		ans = append(ans, nums[nums[idx]])
	}

	return
}
