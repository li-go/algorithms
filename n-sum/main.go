package main

import "fmt"

func main() {
	fmt.Println("threeSum:", threeSum(
		[]int{-1, 0, 1, 2, -1, -4},
	))
	fmt.Println("fourSum:", fourSum(
		[]int{1, 0, -1, 0, -2, 2}, 0,
	))
	fmt.Println("threeSumClosest:", threeSumClosest(
		[]int{0,2,1,-3}, 1,
		//[]int{1, 1, 1, 1}, 0,
	))
}
