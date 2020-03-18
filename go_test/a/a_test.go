package a_test

import "testing"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TestConstantTry(t *testing.T) {
	nums := []int{3, 2, 3}
	target := 6
	t.Log(twoSum(nums, target))
}
