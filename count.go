package main

import (
	"fmt"
	"strings"
)

func wordCount(word string) map[string]int {
	Count := make(map[string]int)
	for _, i := range strings.Fields(word) {
		Count[i]++
	}
	return Count
}
func main() {
	fmt.Println(wordCount("go run me we are go run we are me coming"))
}
