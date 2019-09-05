// https://leetcode.com/problems/generate-parentheses/
package main

func generateParenthesis(n int) []string {
	var result []string
	var dfs func(s string, l, r int)
	dfs = func(s string, l, r int) {
		if l == 0 && r == 0 {
			result = append(result, s)
			return
		}
		if l > 0 {
			dfs(string(append([]uint8(s), '(')), l-1, r)
		}
		if r > l {
			dfs(string(append([]uint8(s), ')')), l, r-1)
		}
	}
	dfs("", n, n)
	return result
}
