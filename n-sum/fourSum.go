// https://leetcode.com/problems/4sum/
package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			twoSums := twoSum_fourSum(nums, j+1, len(nums)-1, target-nums[i]-nums[j])
			for _, s := range twoSums {
				result = append(result, []int{nums[i], nums[j], s[0], s[1]})
			}
		}
	}
	return result
}

func twoSum_fourSum(nums []int, l, r int, target int) [][]int {
	var result [][]int
	for l < r {
		if nums[l]+nums[r] == target {
			result = append(result, []int{nums[l], nums[r]})
			l++
			r--
			// remove duplicate
			for l < r && nums[l] == nums[l-1] {
				l++
			}
			for l < r && nums[r] == nums[r+1] {
				r--
			}
		}
		for l < r && nums[l]+nums[r] < target {
			l++
		}
		for l < r && nums[l]+nums[r] > target {
			r--
		}
	}
	return result
}
