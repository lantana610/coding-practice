package main

import (
	"fmt"
	"strconv"
	"strings"
)

func input() {
	word := ""
	operation := ""
	index := ""
	for {
		fmt.Println("Please Do Input A word")
		fmt.Scan(&word)
		fmt.Println("Please choose an operation: (lastChar or capitalize or deleteIndex): ")
		fmt.Scan(&operation)
		switch operation {
		case "lastChar":
			fmt.Println(lastChar(word))

		case "capitalize":
			fmt.Println(capitalize(word))

		case "deleteIndex":
			fmt.Println("Provide the index to be deleted")
			fmt.Scan(&index)
			n, err := strconv.Atoi(index)
			if err != nil {
				fmt.Println("Provide a valid number")
			}
			if n >= len(word) {
				fmt.Println("Index out of range")
				continue
			}
			fmt.Println(deleteIndex(word, n))
		default:
			fmt.Println("Unknown command use: lastChar || capitalize || deleteIndex")

		}
		break

	}
}

func lastChar(str string) string {
	return str[len(str)-1:]
}
func capitalize(str string) string {
	return strings.ToUpper(str)
}
func deleteIndex(str string, i int) string {
	return str[:i] + str[i+1:]
}
