package main 

import(
	"fmt"
)

func back(word string) string{
	m := ""
	for i :=  len(word)-1;i >= 0; i--{
		m += string(word[i])
	}
	return m
}
func main(){
	fmt.Println(back("palindrome"))
}