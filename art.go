package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	font := make(map[rune][]string)
	data, err := os.ReadFile(filename)

	if len(data) == 0 {
		return nil, errors.New("empty file")
	}
	if err != nil {
		fmt.Println("eror")
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	if len(lines) < 856 {
		return nil, errors.New("invalid content")
	}
	for c := ' '; c <= '~'; c++ {
		start := (int(c) - 32) * 9
		font[c] = lines[start+1 : start+9]

	}
	var result strings.Builder
	input := strings.ReplaceAll(filename, "\\n", "\n")
	word := strings.Split(input, "\n")

	for i, j := range word {
		if j == "" {
			if i != len(word)-1 {
				result.WriteString("\n")

			}
			continue

		}
		for row := 0; row < 8; row++ {
			for _, c := range j {
				result.WriteString(font[c][row])

			}
			result.WriteString("\n")
		}
	}
	return font, err
}
