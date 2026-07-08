package main

import (
	"fmt"
)

func valid(s string) bool {
	return s == "addition" || s == "subtraction" || s == "multiplication" || s == "division" || s == "modulus"
}
func calculate() {
	colour := "\033[34m"
	reset := "\033[0m"
	green := "\033[32m"
	red := "\033[31m"

	fmt.Println(green + "welcome to lanta digital calculator, what calculation can we help you with: " + reset)
	var option string
	fmt.Println("addition")
	fmt.Println("subtraction")
	fmt.Println("division")
	fmt.Println("multiplication")
	fmt.Println("modulus")
	fmt.Println("help")
	for {

		fmt.Println(green + "input an option: " + reset)
		fmt.Scan(&option)

		if option == "help" {
			fmt.Println(colour + "welcome to lanta digital calculator, this is a digital calculator that can help you with the following operation (adition, multiplication, subtraction, division, modulus), of any number of your chioce, all you have to do is to enter any operation of your chioce from the option given input the first and second digit which you want to calculate and it will calculated with the answer for you. thank for using lanta digital calculator!!" + reset)
			continue
		}

		if !valid(option) {
			fmt.Println(red + "invalid operation do you need help go to help menu" + reset)
			continue
		}

		var digit1 int
		fmt.Println(green + "enter digit1" + reset)
		_, err1 := fmt.Scan(&digit1)
		var digit2 int
		fmt.Println(green + "enter digit2" + reset)
		_, err2 := fmt.Scan(&digit2)

		// if digit1 && digit2 != {
		// 	fmt.Println("invalid operation digit must be a number")
		// }

		if err1 != nil && err2 != nil {
			fmt.Println(red + "invalid operation digit must be a number" + reset)
			continue
		}

		if option == "addition" {
			result := digit1 + digit2
			fmt.Printf("%v + %v  =  %v \n", digit1, digit2, result)
		}
		if option == "subtraction" {
			result := digit1 - digit2
			fmt.Printf("%v - %v = %v \n", digit1, digit2, result)
		}
		if option == "division" {

			if digit2 == 0 {
				fmt.Println(red + "cannot divide number by 0" + reset)
				continue
			}

			result := digit1 / digit2
			fmt.Printf("%v / %v = %v \n", digit1, digit2, result)
			break
		}
		if option == "multiplication" {
			result := digit1 * digit2
			fmt.Printf("%v * %v = %v \n", digit1, digit2, result)
		}
		if option == "modulus" {
			result := digit1 % digit2
			fmt.Printf("%v modulus %v = %v \n", digit1, digit2, result)
		}

	}
}
func main() {
	calculate()
}
