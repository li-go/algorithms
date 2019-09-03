// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
package main

var digitToLetters = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	if len(digits) == 1 {
		return digitToLetters[digits[0]]
	}
	var result []string
	ss := letterCombinations(digits[1:])
	for _, c := range digitToLetters[digits[0]] {
		for _, s := range ss {
			result = append(result, c+s)
		}
	}
	return result
}
