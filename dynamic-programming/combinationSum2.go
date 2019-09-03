// https://leetcode.com/problems/combination-sum-ii/
package main

import (
	"sort"
)

type Comb2 struct {
	index []int
	value []int
}

type Comb2s []Comb2

func deepEqual2(a, b []int) bool {
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

func (c *Comb2s) Append(comb Comb2) {
	sort.Ints(comb.value)
	for _, comb1 := range *c {
		if deepEqual2(comb1.value, comb.value) {
			return
		}
	}
	*c = append(*c, comb)
}

func (c *Comb2s) Values() [][]int {
	var result [][]int
	for _, comb := range *c {
		result = append(result, comb.value)
	}
	return result
}

func contains(slice []int, i int) bool {
	for _, x := range slice {
		if x == i {
			return true
		}
	}
	return false
}

func combinationSum2(candidates []int, target int) [][]int {
	dp := make([]Comb2s, target+1)
	for n := 1; n <= target; n++ {
		for i, c := range candidates {
			if n-c == 0 {
				dp[n].Append(Comb2{
					index: []int{i},
					value: []int{c},
				})
			} else if n-c > 0 {
				for _, comb := range dp[n-c] {
					if contains(comb.index, i) {
						continue
					}

					dp[n].Append(Comb2{
						index: append([]int{i}, comb.index...),
						value: append([]int{c}, comb.value...),
					})
				}
			}
		}
	}
	//fmt.Println(dp)
	return dp[target].Values()
}
