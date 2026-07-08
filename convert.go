package main

import (
	"fmt"
)

func kiloconvert() {
	var convert string

	for {
	start:
		fmt.Println("please enter what you want to do")
		fmt.Println("(1)kilogram_to_gram")
		fmt.Println("(2)gram_to_kilogram")
		fmt.Println("(3)help_menu")
		fmt.Scan(&convert)

		if convert != "1" && convert != "2" && convert != "3" {
			fmt.Println("invalid operation do you need help go to help menu")
			goto start
		}

		if convert == "3" {
			fmt.Println("thanks for using my unit converter app here we only covert (kilogram to gram and from gram to kilogram),all u have to do is to choose your convertion method and")
			goto start
		}

		var kilogram int

		var gram float64

		if convert == "1" {
			fmt.Println("enter the the kilogram to convert")
			_, err1 := fmt.Scan(&kilogram)
			if err1 != nil {
				fmt.Println("invalid operation please do input a number")
				continue
			}

			converted := kilogram * 1000
			fmt.Printf("%v kilogram converted to %vgram  \n", kilogram, converted)
		}
		if convert == "2" {
			fmt.Println("enter the gram you want to convert")
			_, err2 := fmt.Scan(&gram)
			if err2 != nil {
				fmt.Println("invalid operation please do input a number")

			}
			converted := gram / 1000.0
			fmt.Printf("%v gram converted %v \n", gram, converted)
		}
		if convert == "3" {
			fmt.Println("welcome to my unit converter app, this is a unit convert that help convert  kilogram to gram and from gram to kilogram, all you have to do is choose from the options given what you want do input the number to convert and it will be done,thank you!!")
			goto start
		}

	}
}

func main() {
	kiloconvert()
}
