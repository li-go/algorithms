// https://leetcode.com/problems/combination-sum/
package main

import (
	"sort"
)

type Combs [][]int

func deepEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (c *Combs) Append(comb []int) {
	sort.Ints(comb)
	for _, comb1 := range *c {
		//if reflect.DeepEqual(comb1, comb) {
		if deepEqual(comb1, comb) {
			return
		}
	}
	*c = append(*c, comb)
}

func combinationSum(candidates []int, target int) [][]int {
	dp := make([]Combs, target+1)
	for n := 1; n <= target; n++ {
		for _, c := range candidates {
			if n-c == 0 {
				dp[n].Append([]int{c})
			} else if n-c > 0 {
				for _, comb := range dp[n-c] {
					dp[n].Append(append([]int{c}, comb...))
				}
			}
		}
	}
	//fmt.Println(dp)
	return dp[target]
}
