package main

import "strings"

func RenderLine(text, color, substring string, banner map[rune][]string) []string {
	var output []string
	start, end := SubstrX(text, substring)
	for i := range 8 {

		var result strings.Builder
		for j, c := range text {
			render := banner[c][i]
			colored := false
			for k := range start {
				if j >= start[k] && j < end[k] {
					colored = true
					break
				}
			}
			if colored {
				result.WriteString(asciicolour(color, render))
			} else {
				result.WriteString(render)
			}
		}
		output = append(output, result.String())
	}
	return output
}
func SubstrX(text, substring string) ([]int, []int) {
	start := []int{}
	end := []int{}
	if substring == "" {
		start, end = []int{0}, []int{len(text)}
		return start, end
	}

	for i := 0; i <= len(text)-len(substring); i++ { // to know the last valid starting position of the substring
		match := true
		for j := 0; j < len(substring); j++ { // to know how many characters to check i.e length of substring
			if text[i+j] != substring[j] {
				match = false
				break
			}
		}
		if match {
			start = append(start, i)
			end = append(end, i+len(substring))

		}

	}
	return start, end

}
