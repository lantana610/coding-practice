package main

import(
	"fmt"
	"strconv"
	"strings"
)
func last(words string)string {
	j := strings.Fields(words)
	for i := 0; i < len(j); i++{
		if j[i] == "(hex)" {
			data, err := strconv.ParseInt(j[i-1], 16, 64)
			if err == nil{
				j[i-1] = strconv.FormatInt(data, 10)
				j = append(j[:i], j[i+1:]...)
				i--

			}
		}
		
	}
	return strings.Join(j, " ")

}
func first(words string)string {
	j := strings.Fields(words)
	for i := 0; i < len(j); i++{
		if j[i] == "(bin)" {
			data, err := strconv.ParseInt(j[i-1], 2, 64)
			if err == nil{
				j[i-1] = strconv.FormatInt(data, 10)
				j = append(j[:i], j[i+1:]...)
				i--

			}
		}
		
	}
	return strings.Join(j, " ")

}
func main(){
	fmt.Println(last("1E (hex) file were added"))
	fmt.Println(first("10 (bin) file were added"))

}