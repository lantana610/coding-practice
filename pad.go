package main

import "strings"

func padArt(str []string, with int) []string {
	result := make([]string, len(str))

	for i, j := range str{
		if with <= 0 || len(j) >= with{
			result[i] = j
			continue
		}
		padding := strings.Repeat(" ", with-len(j))
		result[i] = j + padding
	}
	return result

}
