package main

import (
	"fmt"
)

func FCconvert() {
	var temperature string
start:
	fmt.Println("enter the temperature you want to convert")
	fmt.Println("[1]c-f")
	fmt.Println("[2]f-c")
	fmt.Println("[3]help")
	fmt.Println("[4]end")
	fmt.Scan(&temperature)

	for {
		if temperature != "1" && temperature != "2" && temperature != "3" && temperature != "4" {
			fmt.Println("invalid operation do you need help go help menu")
			goto start
		}
		if temperature == "1" {
			var celsius float64
			fmt.Println("enter what you want to convert")
			_, err1 := fmt.Scan(&celsius)
			if err1 != nil {
				fmt.Println("invalid flaot64 celsius must be a number")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			result := (celsius * 9 / 5) + 32
			fmt.Printf("%v celsius = %v fahrenheit\n", celsius, result)

		}

		if temperature == "2" {
			var fahrenheit float64
			fmt.Println("enter what you want to convert")
			_, err2 := fmt.Scan(&fahrenheit)
			if err2 != nil {
				fmt.Println("invalid float64 fahrenheit must be a number")
				var empty string
				fmt.Scanln(&empty)
				continue
			}
			result := (fahrenheit - 32) * 5 / 9
			fmt.Printf("%v fahrenheit = %v celsius\n", fahrenheit, result)

		}

		if temperature == "3" {
			fmt.Println("this is temperature convertion app, we convert from celsius to fahrenheit and from fahrenheit to celsuis, just choose from the option given which convertion you want to do and input the number you want to convert, thank you!! ")
			goto start
		}

		if temperature == "4" {
			break
		}

	}
}
func main() {
	FCconvert()
}
