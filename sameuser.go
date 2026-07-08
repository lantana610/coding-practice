package main

import (
	"fmt"
	"strings"
)

func userinput() {
	for {
		var word string
		fmt.Print("input a word: ")
		fmt.Scanln(&word)

		var options string
		fmt.Print("choose an operation from the options: (lastchar, capitalize, deleteIndex) ")
		fmt.Scan(&options)

		if options == "lastchar" {
			fmt.Println(word[len(word)-1:])
		}

		if options == "capitalize" {
			fmt.Println(strings.ToUpper(word))
		}

		if options == "deleteIndex" {
			var index int

			fmt.Print("input the index you want to delete: ")
			fmt.Scan(&index)

			if index < 0 && index >= len(word) {

				fmt.Println("index is out of range: ")
				continue
			}
			fmt.Println(word[:index] + word[index+1:])

		}

	}

}

// func main() {
// 	userinput()
// }
