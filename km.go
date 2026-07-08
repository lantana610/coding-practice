package main

import (
	"fmt"
)

func kmconvert() {
	var conversion string
start:
	fmt.Println("what conversion can we help with")
	fmt.Println("[1]kilometer_to_meter")
	fmt.Println("[2]meter_to_kilometer")
	fmt.Println("[3]help_menu")
	fmt.Println("[4]quit")
	fmt.Scan(&conversion)

	for {
		if conversion != "1" && conversion != "2" && conversion != "3" && conversion != "4" {
			fmt.Println("invalid operation do you need help go to help menu")
			goto start
		}
		if conversion == "1" {
			var kilometer int
			fmt.Println("enter the kilometer you want convert")
			_, err1 := fmt.Scan(&kilometer)
			result := kilometer * 1000
			fmt.Printf("%v kilometer = %v meter\n", kilometer, result)
			if err1 != nil {
				fmt.Println("invalid number kilometer must be a number")
				goto start
			}
		}
		if conversion == "2" {
			var meter float64
			fmt.Println("enter the meter you want to convert")
			_, err2 := fmt.Scan(&meter)
			result := meter / 1000.0
			fmt.Printf("%v meter = %v kilometer\n", meter, result)
			if err2 != nil {
				fmt.Println("invalid number meter must be a number")
				goto start
			}
		}
		if conversion == "3" {
			fmt.Println("weelcome to my meter conversion app, this is a meter conversion app that can only help you with converting kilometer to meter and meter to kilometer, so just input what you want to convert from which either from kilometer or from meter, thank you for using this app!! ")
			goto start
		}
		if conversion == "4" {
			break
		}

	}

}
func main() {
	kmconvert()
}
