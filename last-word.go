package main

import(
	"fmt"
	"strings"
)
func las(words []string)[]string{
	for i := 0; i < len(words); i++{
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
		
	}
	return words

}
func me(words []string)[]string{
	for i := 0; i < len(words); i++{
		if words[i] == "(cap)" {
			words[i-1] = strings.ToLower(words[i-1])
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
		
	}
	return words

}
func main(){
	fmt.Println(las([]string{"WORD", "(low)", "ME", "(low)", "ME"}))
	fmt.Println(me([]string{"WORD", "we", "ME",  "ME","(cap)"}))

}