// https://leetcode.com/problems/3sum/
package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		result = append(result, twoSum_threeSum(nums, i+1, len(nums)-1, -nums[i])...)
	}
	return result
}

func twoSum_threeSum(nums []int, l, r int, target int) [][]int {
	var result [][]int
	for l < r {
		if nums[l]+nums[r] == target {
			result = append(result, []int{-target, nums[l], nums[r]})
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
