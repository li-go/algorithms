// https://leetcode.com/problems/3sum-closest/
package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := 1 << 32
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		min := abs(closestSum - target)
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < min {
				min = abs(sum - target)
				closestSum = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return closestSum
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
