package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func BannerLoader(input string) (map[rune][]string, error) {
	data, err := os.ReadFile(input + ".txt")
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	if err != nil {
		fmt.Println("error while reading file")
		os.Exit(1)
	}

	line := strings.Split(string(data), "\n")
	if len(line) > 856 {
		return nil, errors.New("incomplete banner file")
	}
	font := make(map[rune][]string)

	for char := ' '; char <= '~'; char++ {
		begin := (int(char) - 32) * 9
		font[char] = line[begin+1 : begin+9]
	}
	return font, nil
}

func BuildArt(text string, banner map[rune][]string) string {
	if text == "" {
		return ""
	}

	err := ValidInput(text)
	if err != nil {
		return "invalid ascii character"
	}

	word := SplitText(text)

	var result strings.Builder
	for i, char := range word {
		if char == "" {
			if i == len(word)-1 {
				continue
			}
			result.WriteByte('\n')
		}
		output := RenderLine(char, banner)
		for _, value := range output {
			result.WriteString(value)
			result.WriteString("\n")
		}

	}

	return result.String()
}

func RenderLine(input string, banner map[rune][]string) []string {

	result := []string{}
	for i := 0; i <= 7; i++ {
		var output strings.Builder
		for _, ch := range input {
			output.WriteString(banner[ch][i])
		}
		result = append(result, output.String())
	}
	return result
}

func SplitText(str string) []string {
	output := strings.ReplaceAll(str, "\r\n", "\n")
	return strings.Split(output, "\n")

}

func ValidInput(input string) error {

	for _, r := range input {
		if r == '\n' || r == '\r' {
			continue
		}
		if r < 32 || r > 126 {
			return fmt.Errorf("%v is not a valid ascii character", r)
		}
	}
	return nil
}
