package main

import (
	"fmt"
)

func main() {
	fmt.Println("combinationSum:", combinationSum(
		[]int{7, 3, 2}, 18,
	))
	fmt.Println("combinationSum2:", combinationSum2(
		[]int{10, 1, 2, 7, 6, 1, 5}, 8,
	))
}
