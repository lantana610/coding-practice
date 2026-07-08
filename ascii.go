package main

import (
	"fmt"
	"strings"
)

func Ascii(str string, banner [][]string) (string, error) {
	if len(banner) < 1 {
		return "", fmt.Errorf("banner files are empty")

	}

	const Height = 8
	var output strings.Builder

	lines := strings.Split(str, "\n")

	for _, i := range lines {
		for j := 0; j < Height; j++ {
			for _, ch := range i {
				index := int(ch) - 32
				if index < 0 || index >= len(banner) {
					return "", fmt.Errorf("unsupported character: %q", ch)
				}
				output.WriteString(banner[index][j])
			}
			output.WriteString("\n")

		}
	}
	return output.String(), nil
}

