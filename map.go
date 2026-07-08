package main

import "fmt"

func main() {
	var b = map[string]string{"Name": "Lize", "Age": "30", "Nationality": "South African"}
	a := map[string]string{"Brand": "Volkswagen", "Model": "Atlas", "Year": "2027"}

	fmt.Printf("b\t%v\n", b)
	fmt.Printf("a\t%v\n", a)
}
