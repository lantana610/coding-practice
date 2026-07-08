package main

import "strings"

func trim(str []string) []string {
	result := make([]string, len(str))

	for i, j := range str {
		result[i] = strings.TrimRight(j, " ")
	}
	return result
}
