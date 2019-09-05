package main

import (
	"fmt"
)

func main() {
	// https://leetcode.com/problems/similar-string-groups/
	fmt.Println("numSimilarGroups:", numSimilarGroups([]string{
		"kccomwcgcs", "socgcmcwkc", "sgckwcmcoc", "coswcmcgkc", "cowkccmsgc", "cosgmccwkc", "sgmkwcccoc", "coswmccgkc", "kowcccmsgc", "kgcomwcccs",
	}))
}

func numSimilarGroups(A []string) int {
	isSimilar := func(a, b string) bool {
		var diff int
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
				if diff > 2 {
					return false
				}
			}
		}
		return true
	}

	group := make([]int, len(A))
	var groupCount int

	var grouping func(int)
	grouping = func(n int) {
		for i := 0; i < len(A); i++ {
			if group[i] == 0 && isSimilar(A[i], A[n]) {
				group[i] = group[n]
				grouping(i)
			}
		}
	}

	for i := 0; i < len(A); i++ {
		if group[i] == 0 {
			groupCount++
			group[i] = groupCount
			grouping(i)
		}
	}

	return groupCount
}
