// https://leetcode.com/problems/permutations-ii/
package main

import "sort"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	visited := make([]bool, len(nums))
	var dfs func([]int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] || i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			dfs(append(path, nums[i]))
			// backtrack
			visited[i] = false
		}
	}
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	dfs(path)
	return result
}
