package main

import (
	"fmt"
)

func main() {
	fmt.Println("findLadders:", findLadders(
		"hit",
		"cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"},
	))
}
